import http from 'k6/http';
import { check } from 'k6';

export const options = {
  scenarios: {
    read_test: {
      executor: 'constant-vus',
      vus: Number(__ENV.VUS || 10000),
      duration: __ENV.DURATION || '30s',
    },
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:3000';

export default function () {
  const res = http.get(`${BASE_URL}/products?page=1&limit=100`);
  check(res, { 'status 200': r => r.status === 200 });
}
