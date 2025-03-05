/**
 * @package
 */
'use client';

import { RiAddLine, RiCake2Line, RiIdCardLine, RiLockLine, RiMailLine, } from "react-icons/ri";
import { useContext, useEffect, useState } from "react";
import { Controller, useForm, ValidationRule } from "react-hook-form";
import InputField from "@/components/layout/input/inputField";
import { CalloutContext } from "@/app/state";
import { useRouter } from "next/navigation";
import TextWithIcon from "@/components/ui/text/textWithIcon";
import { registerApi } from "./action";
import { CalendarDate, today, getLocalTimeZone } from "@internationalized/date";
import DatePickerField from "@/components/layout/input/datePickerField";
import DPLink from "@/components/ui/text/link";
import DPButton from "@/components/ui/button/button";

type FormContent = {
  email: string;
  password: string;
  name: string;
  birthday: CalendarDate;
  passwordConfirmation: string;
}

export default function UserRegisterPresentation(){
  const {callout, setCallout} = useContext(CalloutContext);
  const router = useRouter();

  const { control, handleSubmit, reset, watch } = useForm<FormContent>({
    defaultValues: {
      email: "",
      password: "",
      name: "",
      birthday: today(getLocalTimeZone()),
      passwordConfirmation: ""
    } 
  });

  const [isSubmit, setIsSubmit] = useState(false);

  useEffect(() => {
    if (isSubmit) {
      const loginFlow = async () => {
        const res = await registerApi(watch().email, watch().password, watch().birthday.toString(), watch().name, watch().passwordConfirmation);

        if (res.errors) {
          setCallout([...callout, {content: res.errors[0], type: 'error'}]);
        }
        else {
          setCallout([...callout, {content: "登録したメールアドレス宛に確認コードを送信しました", type: 'info'}]);
          router.push("/confirm-email");
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
        <TextWithIcon icon={<RiAddLine />} size="text-4xl">
          <h1 className="text-4xl font-bold tracking-widest text-foreground">新規登録</h1>
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
                field={field}
                fieldState={fieldState}
                title="メールアドレス"
                type="email"
                isRequired
                autoFocus
              ></InputField>
            )}
          ></Controller>

          <Controller
            name="name"
            control={control}
            rules={{ required: "ユーザ名が未入力です", max: {
              value: 50,
              message: "ユーザ名は50文字以下で入力してください"
            } }}
            render={({ field, fieldState }) => (
              <InputField
                title="ユーザ名"
                type="text"
                field={field}
                fieldState={fieldState}
                isRequired
                helperText="特殊記号は使用できません"
                icon={<RiIdCardLine />}
              ></InputField>
            )}
          ></Controller>

          <Controller
            name="birthday"
            control={control}
            rules={{ required: "生年月日が未入力です", validate: (value) => value.compare(today(getLocalTimeZone())) < 0 || "未来の日付は指定できません" }}
            render={({ field, fieldState }) => (
              <DatePickerField
                title="生年月日"
                field={field}
                fieldState={fieldState}
                isRequired
                icon={<RiCake2Line />}
              ></DatePickerField>
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

          <Controller
            name="passwordConfirmation"
            control={control}
            rules={{ required: "確認用のパスワードが未入力です", pattern: passwordValidationRule, validate: (value) => value === watch().password || "パスワードが一致しません" }}
            render={({ field, fieldState }) => (
              <InputField
                title="パスワードの確認"
                type="password"
                field={field}
                fieldState={fieldState}
                isRequired
                helperText="パスワードをもう一度入力してください"
                icon={<RiLockLine />}
              ></InputField>
            )}
          ></Controller>

          <hr className="border-subtext" />

          <div className="flex justify-center mt-6">
            <DPButton colormode="secondary" type="submit">
              <TextWithIcon icon={<RiAddLine />}>新規登録</TextWithIcon>
            </DPButton>
          </div>
        </form>

        <div className="flex justify-center">
          <span>
          アカウントをお持ちの方は
          <DPLink href="/login">
            ログイン
          </DPLink>
          から！
          </span>
        </div>
      </main>
    </div>
  )
}