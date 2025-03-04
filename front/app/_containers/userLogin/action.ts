"use server";

import { apiPrefix } from "@/constants/api";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

type ErrorResponse = {
  status: number;
  errors?: string[];
}

export const loginApi = async (email: string, password: string): Promise<ErrorResponse> => {
  const cookie = await cookies();

  const res = await fetch(`${apiPrefix}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email,
      password,
    }),
    credentials: "include",
    cache: "no-cache",
  });
  
  if (res.ok) {
    const data = await res.json();
    cookie.set("devport_api_token", data.Token);

    return { status: res.status };
  } else {
    const error = await res.json();

    return {
      status: res.status,
      errors: error.errors,
    }
  }
}
