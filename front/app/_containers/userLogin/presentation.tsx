/**
 * @package
 */
'use client';

import { RiLockLine, RiMailLine } from "react-icons/ri";
import { useActionState, useContext, useEffect, useRef, useState } from "react";
import { Controller, useForm, ValidationRule } from "react-hook-form";
import InputField from "@/components/ui/input/inputField";
import {loginApi} from "@/app/_containers/userLogin/action";
import { CalloutContext, calloutItemType } from "@/app/state";

type FormContent = {
  email: string;
  password: string;
}

export default function UserLoginPresentation(){
  const {callout, setCallout} = useContext(CalloutContext);

  const { control, handleSubmit, reset, watch,  } = useForm<FormContent>({
    defaultValues: {
      email: "",
      password: "",
    }
  });

  const [isSubmit, setIsSubmit] = useState(false);

  useEffect(() => {
    if (isSubmit) {
      const loginFlow = async () => {
        const res = await loginApi(watch().email, watch().password);

        if (res.errors) setCallout([...callout, {content: res.errors[0], type: 'error'}]);
        else setCallout([...callout, {content: "ログインに成功しました", type: 'success'}]);
      }
      loginFlow();
      reset();
      setIsSubmit(false);
    }
  }, [isSubmit]);

  const passwordValidationRule: ValidationRule<RegExp> = {
    value: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[.+\-[\]*~_#:?]).{8,64}$/,
    message: "パスワードは英数字をそれぞれ1文字以上含む8文字以上で入力してください"
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <header className="mb-6">
        <h1 className="text-4xl font-bold tracking-widest text-foreground">ログイン</h1>
      </header>

      <main className="flex flex-col gap-4 w-1/3">
        <form onSubmit={handleSubmit(() => setIsSubmit(true))}>
          <Controller
            name="email"
            control={control}
            rules={{ required: "メールアドレスが未入力です" }}
            render={({ field, fieldState }) => (
              <InputField
                title="メールアドレス"
                type="email"
                field={field}
                fieldState={fieldState}
                isRequired
              ></InputField>
            )}
          ></Controller>
          <Controller
            name="password"
            control={control}
            rules={{ required: "パスワードが未入力です", pattern: passwordValidationRule }}
            render={({ field, fieldState }) => (
              <InputField
                title="パスワード"
                type="password"
                field={field}
                fieldState={fieldState}
                isRequired
              ></InputField>
            )}
          ></Controller>

          <button type="submit" className="bg-primary text-white p-2 rounded">
            ログイン
          </button>
        </form>
      </main>
    </div>
  )
}