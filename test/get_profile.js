import http from "k6/http";
import { check, fail } from "k6";
import { actionLogin } from "./helper/auth.js";
import { actionGetProfile } from "./helper/profile.js";

export let options = {
  scenarios: {
    get_profile: {
      executor: "ramping-vus",
      startVUs: 0,
      stages: [
        { duration: "30s", target: 50 },
        { duration: "30s", target: 200 },
        { duration: "1m", target: 800 },
        { duration: "1m", target: 1500 },
        { duration: "30s", target: 3000 },
        { duration: "30s", target: 6000 },
        { duration: "1m", target: 10000 },
        { duration: "30s", target: 0 }, // Ramp down
      ],
      exec: "get_profile",
    },
  },
};

export function get_profile(data) {
    const loginResponse = actionLogin();
    // const token = loginResponse.json().token;
    // const getProfile = actionGetProfile(token);
    console.log(loginResponse.body);
}
