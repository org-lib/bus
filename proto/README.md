git config  https.proxy http://127.0.0.1:19180
git config  https.proxy https://127.0.0.1:19180
powerproto build .

#取消git 代理
git config --global --unset http.proxy
git config --global --unset https.proxy
