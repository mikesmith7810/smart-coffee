import { runCoffeeFlow } from './lib/coffee.js'

// Ramp up to 10 users, hold, then ramp down
export const options = {
  stages: [
    { duration: '30s', target: 5 },   // ramp up to 5 users
    { duration: '1m',  target: 10 },  // ramp up to 10 users
    { duration: '2m',  target: 10 },  // hold at 10 users
    { duration: '30s', target: 0 },   // ramp down
  ],
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<500'],
  },
}

export default function () {
  runCoffeeFlow()
}
