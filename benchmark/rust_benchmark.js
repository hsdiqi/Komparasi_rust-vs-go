import http from "k6/http";
import { check } from "k6";
import { Trend, Rate, Counter } from "k6/metrics";
import exec from "k6/execution";

// ======================================================
// ENVIRONMENT VARIABLES
// ======================================================
//
// Contoh:
//
// BASE_URL=http://192.168.1.10:8080
// VUS=10
// METHOD=GET
// FRAMEWORK=gin
// ENDPOINT=/api/items
//
// ======================================================

const BASE_URL = __ENV.BASE_URL || "https://rust-api.blackgecko.my.id";
// const BASE_URL = __ENV.BASE_URL || 'http://127.0.0.1:3001';
const VUS = parseInt(__ENV.VUS || "10");
const METHOD = (__ENV.METHOD || "GET").toUpperCase();
const FRAMEWORK = (__ENV.FRAMEWORK || "actix").toLowerCase();
const ENDPOINT = __ENV.ENDPOINT || "/products";
const LIMIT = __ENV.LIMIT || "10";

// ======================================================
// CUSTOM METRICS
// Hanya measurement phase yang dimasukkan ke sini
// ======================================================

const measurementDuration = new Trend("measurement_http_req_duration", true);

const measurementFailed = new Rate("measurement_http_req_failed");

const measurementRequests = new Counter("measurement_requests");

// ======================================================
// K6 OPTIONS
// ======================================================

export const options = {
  scenarios: {
    // -------------------------------
    // Warm-up
    // -------------------------------
    warmup: {
      executor: "constant-vus",
      vus: VUS,
      duration: "30s",
      gracefulStop: "0s",
      exec: "warmup",

      tags: {
        phase: "warmup",
        framework: FRAMEWORK,
        method: METHOD,
      },
    },

    // -------------------------------
    // Measurement
    // Mulai setelah warm-up selesai
    // -------------------------------
    measurement: {
      executor: "constant-vus",
      vus: VUS,
      duration: "60s",
      startTime: "30s",
      gracefulStop: "0s",
      exec: "measurement",

      tags: {
        phase: "measurement",
        framework: FRAMEWORK,
        method: METHOD,
      },
    },
  },

  thresholds: {
    measurement_http_req_failed: ["rate<0.01"],

    measurement_http_req_duration: ["p(95)<1000"],
  },

  summaryTrendStats: ["avg", "min", "med", "max", "p(90)", "p(95)", "p(99)"],
};

// ======================================================
// DATA PRODUCT
// Dibuat mengikuti pola seed.sql
// ======================================================

const categories = [
  "phone",
  "laptop",
  "tablet",
  "accessory",
  "camera",
  "audio",
  "monitor",
  "keyboard",
  "mouse",
  "printer",
];

// ======================================================
// GENERATE PAYLOAD POST
// ======================================================

function generateProduct() {
  const id = __VU * 1000000 + __ITER;

  return JSON.stringify({
    name: `Product ${id}`,
    description: `Description for product ${id}`,
    price: ((id % 1000) + 1) * 1000,
    stock: id % 500,
    category: categories[id % categories.length],
  });
}

// ======================================================
// HTTP PARAMS
// ======================================================

const params = {
  headers: {
    "Content-Type": "application/json",
  },
};

// ======================================================
// TIMESTAMP
// ======================================================

function formatTimestamp() {
  const now = new Date();

  return now.toLocaleString("sv-SE", {
    timeZone: "Asia/Jakarta",
  });
}

export function setup() {
  const startTime = formatTimestamp();

  console.log(`=== TEST START === ${startTime}`);
  console.log(`=== FRAMEWORK === ${FRAMEWORK}`);
  console.log(`=== METHOD === ${METHOD}`);

  console.log(`=== VUS === ${VUS}`);
  console.log(`=== ENDPOINT === ${ENDPOINT}`);
  return {
    startTime,
  };
}

function sendRequest() {
  let url = `${BASE_URL}${ENDPOINT}`;

  let response;

  if (METHOD === "POST") {
    const payload = generateProduct();

    response = http.post(url, payload, {
      ...params,
      tags: {
        endpoint_type: "POST",
        framework: FRAMEWORK,
      },
    });
  } else {
    url = `${url}?limit=${LIMIT}`;

    response = http.get(url, {
      tags: {
        endpoint_type: "GET",
        framework: FRAMEWORK,
      },
    });
  }

  check(response, {
    "status is successful": (r) =>
      r.status >= 200 && r.status < 300,
  });

  return response;
}

// ======================================================
// WARM-UP
// Tidak dimasukkan ke custom measurement metrics
// ======================================================

export function warmup() {
  sendRequest();
}

// ======================================================
// MEASUREMENT
// Data utama penelitian
// ======================================================

export function measurement() {
  if (exec.scenario.iterationInTest === 0) {
    console.log(`=== MEASUREMENT START === ${formatTimestamp()}`);
  }

  const response = sendRequest();

  measurementDuration.add(response.timings.duration);

  measurementFailed.add(response.status < 200 || response.status >= 300);

  measurementRequests.add(1);
}

// export function measurement() {
//   const response = sendRequest();

//   measurementDuration.add(response.timings.duration);

//   measurementFailed.add(
//     response.status < 200 || response.status >= 300
//   );

//   measurementRequests.add(1);
// }

export function teardown(data) {
  const endTime = formatTimestamp();

  //   console.log(`=== TEST END === ${endTime}`);
  console.log(`=== TEST START === ${data.startTime}`);
  console.log(`=== TEST END === ${endTime}`);
}
