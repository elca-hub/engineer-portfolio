/**
 * @package
 */
'use client';

import { RiLockLine, RiMailLine } from "react-icons/ri";
import { useContext, useEffect, useState } from "react";
import { Controller, useForm, ValidationRule } from "react-hook-form";
import InputField from "@/components/layout/input/inputField";
import {loginApi} from "@/app/_containers/login/action";
import { CalloutContext } from "@/app/state";
import { useRouter } from "next/navigation";
import TextWithIcon from "@/components/ui/text/textWithIcon";
import { ButtonStyle } from "@/constants/tailwindConstant";
import DPLink from "@/components/ui/text/link";
import DPButton from "@/components/ui/button/button";

type FormContent = {
  email: string;
  password: string;
}

export default function UserLoginPresentation(){
  const {callout, setCallout} = useContext(CalloutContext);
  const router = useRouter();

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

        if (res.errors) {
          setCallout([...callout, {content: res.errors[0], type: 'error'}]);
        }
        else {
          setCallout([...callout, {content: "ログインに成功しました", type: 'success'}]);
          router.push("/");
        }
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
        <TextWithIcon icon={<RiLockLine />} size="text-4xl">
          <h1 className="text-4xl font-bold tracking-widest text-foreground">ログイン</h1>
        </TextWithIcon>
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
                autoFocus
                icon={<RiMailLine />}
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
                helperText="英数字をそれぞれ1文字以上含む8文字以上で入力してください"
                icon={<RiLockLine />}
              ></InputField>
            )}
          ></Controller>

          <div className="flex justify-center mt-6">
            <DPButton colormode="primary" type="submit">
              <TextWithIcon icon={<RiLockLine />}>ログイン</TextWithIcon>
            </DPButton>
          </div>
        </form>

        <div className="flex justify-center">
          <span>
          アカウントをお持ちでない方は
          <DPLink href="/register">新規登録</DPLink>
          から！
          </span>
        </div>
      </main>
    </div>
  )
}