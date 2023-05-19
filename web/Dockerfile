# 開発環境用
FROM node:18 AS base
CMD [ "bash" ]

# ビルド
FROM base AS test-and-build
COPY . /workspace
# ビルド時に引数を与える（Actions上で指定する）。この引数はビルド時にしか使用しない。
ARG REACT_APP_API_SERVER
WORKDIR /workspace
# npm ci: npm installと似たようなコマンド。厳密な管理が可能
RUN npm ci && npm run test && npm run build

# Nginxをホストする
FROM nginx:latest AS production
# /var/wwwにあるものをHTTPサーバーとして公開する
COPY --from=test-and-build /workspace/build /var/www
COPY nginx.conf /etc/nginx/conf.d/default.conf


