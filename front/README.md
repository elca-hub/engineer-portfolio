# DevPort | front

## 参考文献

1. [Next.jsの考え方 | Zenn](https://zenn.dev/akfm/books/nextjs-basic-principle/viewer/intro)
2. [Next.jsのディレクトリ構成 | Zenn](https://zenn.dev/yutabeee/articles/0f7e8e2fa03946)
3. [Container/Presentationalパターン再入門 | Zenn](https://zenn.dev/buyselltech/articles/9460c75b7cd8d1)

※ ほとんど1.の要約になります。

## ServerとClientをどこで分けるか

Next.jsの最大の強みである「SSR（サーバサイドレンダリング）」。
これを最大限活かすには、Server ComponentとClient Componentをしっかりと理解して、区別する必要があります。

### Client Componentの大きなデメリット

通常のHTMLやJavaScriptは、webサーバから送られてきたhtmlファイルやjsファイルをwebブラウザが解釈し、ユーザに表示します。
つまり、jsの処理は全てクライアント側（webブラウザ）が処理するということになります。
例えばユーザがボタンを押したらカウントアップするといった処理は、当然webブラウザで処理しています。

ここで問題になってくるのが、APIを叩く時です。
「APIを叩く」=「外部サーバにデータをリクエストし、レスポンスされたデータを解釈して処理する」です。
先ほども書いた通り、jsはwebブラウザ側で処理します。
そのため、APIのリクエストはwebブラウザで行われ、レスポンスされたデータもwebブラウザまで渡す必要があります。
大抵の場合、webブラウザとAPIサーバ間の通信には時間がかかりますし、回数が増えるほどそれが顕著になります。
回数を増やせばChatty API（おしゃべりなAPI）になり、回数を減らせばGod APIとなります。
回数を減らした方が確かに速くなりますが、その分設計は複雑化します。
トレードオフということです。

### Server Component

この問題を解決するため、Server Component（正式には「React Server Component」）が生まれました。
これは、Next.jsサーバ側でAPIサーバを叩くという手法です。
こうすることで、すでにAPIで得られた結果を反映したhtmlファイルをwebブラウザ側に渡すことができます。
さらに、APIサーバを公開にする必要がないので、よりセキュアにAPIを取り扱うことができるというメリットもあります。
あとはSEO向上などなど、、、

しかしServer Componentにもデメリットが存在します。
まず一番開発者が感じるデメリットは、従来のClient Componentとは異なる設計手法にする必要があるということ。
イベントハンドリング（`onClick`処理など）ができないため、開発者はディレクトリやファイルの構成を強く意識する必要があります。
そしてあまり開発者が意識しないところだと、RSC Payloadの問題もあります。
これはNext.jsのServer Componentのレンダリングで生じる問題です。

> On the server, Next.js uses React's APIs to orchestrate rendering. The rendering work is split into chunks: by individual route segments and Suspense Boundaries.
> Each chunk is rendered in two steps:
>
> 1. React renders Server Components into a special data format called the React Server Component Payload (RSC Payload).
> 2. Next.js uses the RSC Payload and Client Component JavaScript instructions to render HTML on the server.
>
> Then, on the client:
>
> 1. The HTML is used to immediately show a fast non-interactive preview of the route - this is for the initial page load only.
> 2. The React Server Components Payload is used to reconcile the Client and Server Component trees, and update the DOM.
> 3. The JavaScript instructions are used to hydrate Client Components and make the application interactive.
>
> 引用元：[How are Server Components rendered?](https://nextjs.org/docs/app/building-your-application/rendering/server-components#how-are-server-components-rendered)

翻訳は各自で行ってもらいたいのですが、ここで言いたいのは、Server Componentを使用するとRSC Payloadが生成されるということです。
webブラウザはRSC Payloadを使用してDOMを更新します。
もしもRSC Payloadのサイズが大きければ、せっかくのSSRのメリットを失いかねません。
かといってClient Componentを多くしてしまうと、JavaScriptバンドルのサイズが大きくなってしまいます。
ここでもトレードオフの関係が出てきました。

Server Componentは使い方によってはむしろ速度が遅くなったり、想定より遅い処理になってしまうことがあります。
それをカバーするためにも、Client Componentが必要なのです。

### 暗黙的なClient Component

ディレクトリ構成について語る前に注意したいのが、`use client`の存在です。
通常は`use client`を利用することでClient Componentになりますが、もし`use client`を記述したモジュールをimportした場合、それ以降のモジュールは全て暗黙的にClient Componentと化します。
これを**Client Boundary**と呼びます。
例えば`layout.tsx`に`use client`を書くと、全てのモジュールがClient Componentになります。
これを防ぐためにも、**`use client`をなるべく末端に寄せる**必要があります。

## 設計パターン

### Compositionパターン

大前提として、**Client ComponentからServer Componentを呼び出すことはできません**。
ここで使うのが以下のような形式。

side-menu.tsx

```tsx
'use client'

import { useState } from 'react'

// `children`に`<UserInfo>`などのServer Componentsを渡すことが可能！
export function SideMenu({ children }: { children: React.ReactNode }) {
	const [open, setOpen] = useState(false)

	return (
		<>
			{children}
			<div>
				<button type="button" onClick={() => setOpen((prev) => !prev)}>
					toggle
				</button>
				<div>...</div>
			</div>
		</>
	)
}
```

page.tsx

```tsx
import { UserInfo } from './user-info' // Server Components
import { SideMenu } from './side-menu' // Client Components

/**
 * Client Components(`<SideMenu>`)の子要素として
 * Server Components(`<UserInfo>`)を渡せる
 */
export function Page() {
	return (
		<div>
			<SideMenu>
				<UserInfo />
			</SideMenu>
			<main>{/* ... */}</main>
		</div>
	)
}
```

`SideMenu`はClient Componentですが、`UserInfo`はServer Componentです。
`SideMenu`に`children`を設けることで、Server Componentのモジュールを使用することができます。
これを**Compositionパターン**といいます。

しかしこのパターンを導入する際には、**Server Component**を先に設計しないと後戻りや修正が増える可能性がある\*\*という点には留意が必要です。

### Container/Presentationalパターン

Next.jsでは、Server Componentに対するテストは現状未対応もしくはexperimentalとなっています。
Server ComponentとClient Componentをしっかりと分けて、それぞれにテストを行えるようにする必要があります。
そこで出てくるのが、**Container/Presentationalパターン**です。

このパターンでは、Container ComponentとPresentational Componentの2種類が登場します。
それぞれの責務は以下のとおりです。

|        Component         |                     責務                     |
| :----------------------: | :------------------------------------------: |
|   Container Component    | アプリケーションのロジックに対して責務がある |
| Presentational Component |             UIに対して責務がある             |

Presentational ComponentはUIにしか関心がないため、原則として**状態を持ちません**。
データの受け取り方法はあくまでPropsで行います。

このことから、依存関係は「**Presentational Component -> Container Component**」となります。

このパターン自体はFlux全盛期に登場したようですが、Next.jsにおいてはどうなるのでしょうか？

まずContainer Componentはサーバサイドの処理のみを行います。

Presentational ComponentはClient Componentもしくは**Shared Component**にあたります。

Shared Componentとは、サーバモジュールに依存せず、`use client`の記載もないコンポーネントのことを指します。
Client Boundary内ではClient Componentとして、そうでない場合はServer Componentとして扱われます。

## ディレクトリ構成について

今回はContainer/Presentationalパターンを採用します。

### 設計手法

1. Container Componentを実装する
2. Presentational Componentを実装する

具体例は[こちら](https://zenn.dev/akfm/books/nextjs-basic-principle/viewer/part_2_container_1st_design#%E5%AE%9F%E8%A3%85%E4%BE%8B)を参照してください。

### 今回のディレクトリ構成

以下のようなディレクトリ構成を採用しました。

```text
front/
├── app/
│   ├── parts
│   │   └── [このページでしか利用しないコンポーネント]
│   ├── page.tsx
│   ├── layout.tsx
│   ├── globals.css
│   ├── favicon.ico
│   └── portfolio
│       └── page.tsx
├── public/
│   ├── logo.webp
│   └── [その他画像達]
├── components/
│   ├── layout
│   │   ├── footer
│   │   │   └── footer.tsx
│   │   └── header
│   │       └── header.tsx
│   └── ui
│       ├── button
│       │   ├── button.tsx
│       │   └── buttonIcon.tsx
│       └── text
│           ├── budouxText.tsx
│           └── heading.tsx
├── containers/
│   └── [containerの名称(Ex. userProfile)]
│       ├── index.tsx
│       ├── container.tsx
│       └── presentational.tsx
├── next.config.js
├── package.json
├── tsconfig.json
└── ...
```

### 注意点

まず1つ目に、appフォルダの扱い方です。

基本的にappフォルダ内に様々なフォルダを展開するのがポピュラーかもしれません。
しかしNext.jsはappフォルダをそのままルーティングとして取り扱います。
例えばcomponentsフォルダをappフォルダ内に入れてしまうと、`[ドメイン]/components/...`となってしまいます。
それを解決するために、`_components`とすることで無効化することもできます。
しかし、**ルールは極力少ない方が良い**という個人的なポリシー[^1]によって、特に`_`を気にしなくても大丈夫なようにappフォルダ外に出しました。

2つ目は、partsディレクトリの扱い方です。

partsディレクトリは、そのディレクトリと同じ階層にある`page.tsx`でしか利用されないコンポーネントをまとめます。
もちろん`page.tsx`内に定義することもできますが、そうすると1つのファイルに対するコード量が肥大になり、可読性が低下します。
ファイル単位で明確にコンポーネント化していることを示すためにも、ファイル単位でまとめました。

3つ目は、componentsの中身についてです。

Atomicデザインを適用しても良いのですが、正直あそこまで細かく分けるのはむしろ開発速度が遅くなってしまう可能性があるため、不採用としました。
`ui`ディレクトリには、`button`や`text`など、基本的なコンポーネントをまとめます。
このコンポーネントは**ドメインに一切関与しません（=どこでも利用できる）**。
また、`button`や`text`にも複数のコンポーネントが存在する場合があるため、更にフォルダを分けて細分化しています。

`layout`ディレクトリには、横断的なレイアウトをまとめます。
このコンポーネントも**ドメインに一切関与しません**。
DevPortではheader内にページごとに異なる操作（「ポートフォリオを編集する」、「公開設定」など）があります。
これらの操作とその処理はheader内には記載しません。

`pages`ディレクトリは私独自につけたディレクトリです。
このディレクトリには1つのコンポーネントにしか利用されないコンポーネントを格納します。

[^1]: ルールが多すぎると開発する際に気にすることが増えてストレスになったり、チーム開発においてはメンバーへの共有事項が増えてしまうというデメリットが存在するためです。
