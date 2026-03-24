import { runCoffeeFlow } from './lib/coffee.js'

// Push the service hard — ramp up aggressively, spike, then recover
export const options = {
  stages: [
    { duration: '30s', target: 10 },  // ramp up
    { duration: '1m',  target: 25 },  // increase load
    { duration: '30s', target: 50 },  // spike
    { duration: '1m',  target: 50 },  // hold spike
    { duration: '30s', target: 10 },  // recover
    { duration: '30s', target: 0 },   // ramp down
  ],
  thresholds: {
    http_req_failed: ['rate<0.05'],
    http_req_duration: ['p(95)<2000'],
  },
}

export default function () {
  runCoffeeFlow()
}
