import csv
import sys
import time
from datetime import datetime
from pathlib import Path
from zoneinfo import ZoneInfo

if len(sys.argv) != 3:
    sys.exit(
        "Cara pakai: sudo python3 monitor_resources.py PID hasil.csv"
    )

pid = int(sys.argv[1])
output = sys.argv[2]
proc = Path(f"/proc/{pid}")
timezone = ZoneInfo("Asia/Jakarta")


def process_identity():
    text = (proc / "stat").read_text()
    fields = text.rsplit(")", 1)[1].split()
    return int(fields[19])  # Waktu mulai proses; mendeteksi PID dipakai ulang.


def find_cpu_stat():
    for line in (proc / "cgroup").read_text().splitlines():
        if line.startswith("0::"):
            relative_path = line.split(":", 2)[2]
            return (
                Path("/sys/fs/cgroup")
                / relative_path.lstrip("/")
                / "cpu.stat"
            )

    raise RuntimeError("Cgroup v2 tidak ditemukan")


def read_cpu_usage_usec(cpu_stat):
    for line in cpu_stat.read_text().splitlines():
        key, value = line.split()
        if key == "usage_usec":
            return int(value)

    raise RuntimeError("usage_usec tidak ditemukan")


def read_pss():
    for line in (proc / "smaps_rollup").read_text().splitlines():
        if line.startswith("Pss:"):
            return int(line.split()[1]) / 1024  # KiB → MiB

    raise RuntimeError("Pss tidak ditemukan")


def timestamp():
    return datetime.now(timezone).isoformat(timespec="milliseconds")


def monitor():
    identity = process_identity()
    cpu_stat = find_cpu_stat()

    # Pastikan kedua sumber bisa dibaca sebelum membuat CSV.
    previous_usage = read_cpu_usage_usec(cpu_stat)
    read_pss()
    previous_mono = time.monotonic()
    previous_stamp = timestamp()

    with open(output, "x", newline="") as file:
        writer = csv.writer(file)
        writer.writerow([
            "interval_start",
            "interval_end",
            "interval_seconds",
            "cpu_percent",
            "pss_timestamp",
            "pss_mib",
        ])
        file.flush()

        print(f"Merekam PID {pid} ke {output}", flush=True)
        print("Tekan Ctrl+C setelah pengujian selesai.", flush=True)

        while True:
            time.sleep(1)

            if process_identity() != identity:
                raise RuntimeError(
                    "PID berubah proses. Ambil PID baru dan ulangi monitor."
                )

            current_usage = read_cpu_usage_usec(cpu_stat)
            current_mono = time.monotonic()
            current_stamp = timestamp()

            elapsed = current_mono - previous_mono
            delta_usage = current_usage - previous_usage

            if delta_usage < 0:
                raise RuntimeError("Penghitung CPU direset")

            cpu = delta_usage / 1_000_000 / elapsed * 100
            pss = read_pss()
            pss_stamp = timestamp()

            writer.writerow([
                previous_stamp,
                current_stamp,
                f"{elapsed:.6f}",
                f"{cpu:.4f}",
                pss_stamp,
                f"{pss:.4f}",
            ])
            file.flush()

            print(
                f"{current_stamp} | CPU {cpu:.2f}% "
                f"| PSS {pss:.2f} MiB",
                flush=True,
            )

            previous_usage = current_usage
            previous_mono = current_mono
            previous_stamp = current_stamp


try:
    monitor()
except KeyboardInterrupt:
    print("\nPencatatan selesai.")
except (OSError, RuntimeError, ValueError) as error:
    sys.exit(f"Error: {error}")