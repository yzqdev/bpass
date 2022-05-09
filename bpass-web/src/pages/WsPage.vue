<template>
<div>
  <div class="container">
    <div class="list-group" id="divShow"></div>
    <div>
      <article style="text-align: center">{{info}}</article>
      <div><input class="form-control" v-model="msg" autofocus rows="6" placeholder="请输入发送内容"></div>
      <div><button class="btn btn-default" @click="btnSend" style="margin-top:15px">发 送</button></div>
    </div>
  </div>
</div>
</template>

<script setup lang="ts">
import {useQuasar} from "quasar";
let url='ws://127.0.0.1:8901/ws'
let ws  = new WebSocket(url);
try {
  // ws连接成功
  ws.onopen = function () {
    showInfo("WebSocket Server [" + url +"] 连接成功！");
  };
  // ws连接关闭
  ws.onclose = function () {
    if (ws) {
      ws.close();
      ws = null;
    }
    showError("WebSocket Server [" + url +"] 连接关闭！");
  };
  // ws连接错误
  ws.onerror = function () {
    if (ws) {
      ws.close();
      ws = null;
    }
    showError("WebSocket Server [" + url +"] 连接关闭！");
  };
  // ws数据返回处理
  ws.onmessage = function (result) {
    showWaring(" > " + result.data);
  };
} catch (e) {
  alert(e.message);
}
let msg=$ref('hello')
let info=$ref('hello')
function btnSend() {
  if (ws == null) {
    showError("WebSocket Server [" + url +"] 连接失败，请F5刷新页面!");
    return;
  }
  var content = msg;
  if (content.length <= 0) {
    alert("请输入发送内容!");
    return;
  }

  showSuccess(content);
  ws.send(content);
}
function showInfo(content) {
info=content
}
// 显示警告信息
function showWaring(content) {
  info=content
}
// 显示成功信息
function showSuccess(content) {
  info=content
}
// 显示错误信息
function showError(content) {
  info=content
}
</script>

<style scoped>

</style>
