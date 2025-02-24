# ベースイメージを指定
FROM node:23

# 作業ディレクトリを設定
WORKDIR /app

ADD . .

RUN npm install

CMD ["npm", "run", "dev"]
