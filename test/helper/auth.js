import { check } from "k6";
import http from "k6/http";

export function actionLogin() {
  const url = "http://localhost:8080/api/v1/auth";
  const payload = JSON.stringify({
    email: "anu@gmail.com",
    password: "rahasia123",
    action: "login",
  });
  const params = { headers: { "Content-Type": "application/json" } };

  const response = http.post(url, payload, params);
  if (response.status !== 200) {
    console.log(response.body);
  }

  check(response, {
    "status is 200": (r) => r.status === 200,
  });
  return response;
}
