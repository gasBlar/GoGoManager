import { check } from "k6";
import http from "k6/http";

export function actionGetProfile(token) {
  const url = "http://localhost:8080/api/v1/user";
  const params = { headers: { "Content-Type": "application/json", "Authorization" : "Bearer "+ token} };

  const response = http.get(url, params);
  if (response.status !== 200) {
    console.log(response.body);
  }

  check(response, {
    "status is 200": (r) => r.status === 200,
  });
  return response;
}
