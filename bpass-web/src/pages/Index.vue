<template>
  <q-page>
    <q-tabs
      v-model="tab"
      inline-label
      class="bg-purple text-white shadow-2"
    >
      <q-tab name="mails" icon="mail" label="首页">

      </q-tab>
      <q-tab name="alarms" icon="alarm" label="传输"/>
      <q-tab name="movies" icon="movie" label="扫码"/>
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
        <article class="file-card-container row">
          <q-card class="file-card col-xs-12 col-sm-6 col-md-4" v-for="(item,index) in fileList">
            <q-card-section>
              {{ item.name }}
              <q-separator/>
              <article class="img-container" v-if="item.type=='img'"><img :src="`${baseUrl}/files${item.path}`"
                                                                          :alt="item.name"/></article>
              <article class="img-container" v-else>
                <div> {{ item.ext }}</div>
              </article>
              <article class="file-footer"><span>{{ item.ext }}</span><span>{{ item.sizes }}</span> <span
                class="delete-btn"><q-btn
                style="background:red;color: white" @click="deleteFile(item)">删除</q-btn></span></article>
            </q-card-section>
          </q-card>
        </article>
      </q-tab-panel>

      <q-tab-panel name="alarms">
        <div class="text-h6">Alarms</div>
        Lorem ipsum dolor sit amet consectetur adipisicing elit.
      </q-tab-panel>

      <q-tab-panel name="movies">
        <div class="text-h6">Movies</div>
        Lorem ipsum dolor sit amet consectetur adipisicing elit.
      </q-tab-panel>
    </q-tab-panels>
  </q-page>
</template>

<script>
import {defineComponent} from 'vue';
import {api} from "boot/axios";
import {useQuasar} from 'quasar'
import {baseUrl} from "boot/axios";

export default defineComponent({
  name: 'PageIndex',
  data() {
    return {
      tab: 'mails',
      ips: [],
      pathRoot: '',
      fileList: [],
      baseUrl
    }
  }, methods: {
    openUrl(url) {

      api.get('/api/openurl?url=' + url).then(({data}) => {
        useQuasar().notify("在主电脑打开目录成功!")
      })
    },
    getFileList() {
      api.get("/fileList").then(({data}) => {
        console.log(data)
        this.pathRoot = data.pathRoot
        this.ips = data.ips
        this.fileList = data.fileList
      })
    },
    deleteFile(item) {
      api.get('/api/delete?f=' + item.path).then((data) => {
        this.getFileList()
      })
    }
  }, created() {

    this.getFileList()
  }
})
</script>
<style lang="scss">
.file-card-container {

  .file-card {
    .img-container {
      margin: 1rem;
      display: flex;
      justify-content: center;
      align-items: center;


      div {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 8rem;
        font-size: 2rem;
      }


      img {
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
