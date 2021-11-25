<template>
  <q-page>
    <q-tabs
      v-model="tab"
      inline-label
      class="bg-purple text-white shadow-2"
    >
      <q-route-tab name="mails" icon="mail" label="首页" to="/home">

      </q-route-tab>
      <q-route-tab name="alarms" icon="alarm" label="上传文件" to="/transfer"/>
      <q-route-tab name="text" icon="alarm" label='传输文本' to="/text"/>
      <q-route-tab name="movies" icon="movie" label="聊天" to="/chat"/>
    </q-tabs>
    <q-separator/>

    <q-tab-panels v-model="tab" animated>
      <q-tab-panel name="mails">
        <q-card>
          <q-card-section class="flex flex-center " style="flex-direction: column">

            <article>手机电脑文件传输||局域网文件服务器</article>
            <article v-for="(item,index) in ips">
              <q-btn flat no-caps>{{ item }}</q-btn>
            </article>
            <q-btn icon="folder" no-caps flat @click="openUrl(pathRoot)">{{ pathRoot }}</q-btn>


          </q-card-section>
        </q-card>
        <p></p>
        <q-space/>
        <article v-if="inDir">
          <q-btn color="red" @click="back">返回</q-btn>
        </article>
        <article class="file-card-container row">
          <q-card class="file-card col-xs-12 col-sm-6 col-md-4" v-for="(item,index) in fileList">
            <q-card-section>
              {{ item.name }}
              <q-separator/>
              <article class="img-container" v-if="item.type=='img'"><img :src="`${baseUrl}/files${item.path}`"
                                                                          :alt="item.name"/></article>

              <article class="img-container" v-else>
                <div v-if="item.ext=='dir'" @click="showDir(item)" class="dir-div"> {{ item.ext }}</div>
                <div v-else> {{ item.ext }}</div>
              </article>
              <article class="file-footer"><span>{{ item.ext }}</span><span>{{ item.sizes }}</span> <span
                class="delete-btn"><q-btn
                style="background:red;color: white" @click="deleteFile(item)">删除</q-btn></span></article>
            </q-card-section>
          </q-card>
        </article>
      </q-tab-panel>

      <q-tab-panel name="alarms">
        <q-card>
          <q-card-section>
            <article>
              <q-input style="max-width: 30rem" v-model="text" label="自定义上传子目录" counter dense>
                <template v-slot:prepend>
                  /
                </template>
                <template v-slot:append>
                  /
                </template>


              </q-input>
              <q-uploader
                style="max-width: 30rem;max-height: 50rem;"
                :url="`http://localhost:8901/api/upload?path=${text}`"
                @failed="showFail"
                label="批量上传"
                multiple
                fieldName="file"
                batch

              />
            </article>
          </q-card-section>
        </q-card>
      </q-tab-panel>
      <q-tab-panel name="text">
        <q-card>
          <q-card-section>
            <q-input
              v-model="wsText"
              filled
              @mouseout="setText"
              type="textarea"
            />
          </q-card-section>
        </q-card>
      </q-tab-panel>
      <q-tab-panel name="movies">
        聊天
      </q-tab-panel>
    </q-tab-panels>
  </q-page>
</template>

<script>
import {defineComponent} from 'vue';
import {api} from "boot/axios";
import {useQuasar} from 'quasar'
import {baseUrl} from "boot/axios";

let ws = new WebSocket("ws://" + baseUrl.replace("http://", "") + "/sync/web-socket");
export default defineComponent({
  name: 'PageIndex',
  data() {
    return {
      tab: 'mails',
      ips: [],
      pathRoot: '',
      fileList: [],
      baseUrl, text: '', wsText: '', ws: ws, inDir: false, inPath: ''
    }
  },
  watch: {
    "$route": function (val) {
      if (val) {
        this.getFileList()
      }
    }
  },
  methods: {
    setText() {
      console.log("移出文本框")
      api.post('/api/textdata', {data: this.wsText}).then(({data}) => {
        console.log("发送数据", data)
        this.syncSend("reload_text")
      })
    },
    syncSend(data) {
      console.log("syncsend")
      this.ws.send(data);
    },
    openUrl(url) {
      api.get('/api/openurl?url=' + url).then(({data}) => {
        useQuasar().notify("在主电脑打开目录成功!")
      })
    },
    showFail(err) {
      console.log(err)
    },

    getFileList() {
      api.get("/fileList").then(({data}) => {
        this.pathRoot = data.pathRoot
        this.ips = data.ips
        this.fileList = data.fileList
      })
    },
    deleteFile(item) {
      api.get('/api/delete?f=' + item.path).then((data) => {
        this.getFileList()
      })
    },
    showDir(item) {
      this.inDir = true

      api.get("/fileList?path=" + item.path).then(({data}) => {
        this.pathRoot = data.pathRoot
        this.ips = data.ips
        this.fileList = data.fileList
        this.inPath = item.path
      })
    },
    back() {
      this.inDir = false
      api.get("/fileList").then(({data}) => {
        this.pathRoot = data.pathRoot
        this.ips = data.ips
        this.fileList = data.fileList
      })
    },
    syncDo(data) {
      let msg = data.msg;
      console.log("[syncDo]接受信息", data);
      if (msg === 'reload_text') {
        api.get("/api/textdata",
        ).then((result) => {
          console.log("接受消息成功 ", result)
          this.wsText = result.data;
        }).catch((err) => {
          console.error(err)
        });
      }
    }
  }, created() {

    this.getFileList()
  },
  mounted() {
    api.get("/api/textdata").then((result) => {
      console.log("mounted:textdata:" + JSON.stringify(result.data));
      this.wsText = result.data;
    }).catch((err) => {
      console.error(err)
    });

    let that = this

    this.ws.onmessage = function (result) {
      let data = JSON.parse(result.data);
      console.log("onmessage 0", data)
      //消息接收由载入页面实现
      that.syncDo(data)
    };

  }
})
</script>
<style lang="scss">
.file-card-container {

  .file-card {
    .img-container {
      margin: 1rem;
      height: 8rem;
      display: flex;
      justify-content: center;
      align-items: center;

      .dir-div {
        cursor: pointer;
        border-radius: 0.3rem;
        width: 4rem;
        height: 2rem;
        padding: 0.4rem 0.8rem;
        font-size: 2rem;
        background: #1976D2;
      }

      div {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 8rem;
        font-size: 2rem;
      }


      img {
        cursor: pointer;
        height: 8rem;
      }
    }

    .file-footer {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      align-items: center;

      .delete-btn {
        text-align: right;
      }
    }

  }

}
</style>
