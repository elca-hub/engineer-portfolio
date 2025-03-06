"use server";

import { apiPrefix } from "@/constants/api";
import { cookieConfirmName } from "@/constants/cookieTokenName";
import { makeJwt } from "@/lib/auth/jwt";
import { cookies } from "next/headers";

type ErrorResponse = {
  status: number;
  errors?: string[];
}

export const registerApi = async (email: string, password: string, birthday: string, name: string, passwordConfirmation: string): Promise<ErrorResponse> => {
  const res = await fetch(`${apiPrefix}/signup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email,
      password,
      birthday,
      name,
      "password_confirmation": passwordConfirmation,
    }),
    credentials: "include",
    cache: "no-cache",
  });
  
  if (res.ok) {
    const data = await res.json();
    
    const userEmail = data.Email;

    if (userEmail === undefined) throw new Error("Email is not found in response");

    const payload = {
      email: userEmail as string,
    }

    const token = await makeJwt(payload);

    const cookie = await cookies();

    cookie.set(cookieConfirmName, token, {
      maxAge: 60 * 60 * 24, // 1 day
      path: "/",
    });

    return { status: res.status };
  } else {
    const error = await res.json();

    return {
      status: res.status,
      errors: error.errors,
    }
  }
}
