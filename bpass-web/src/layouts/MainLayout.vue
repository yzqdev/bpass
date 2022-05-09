<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title> bpass </q-toolbar-title>

        <div @click="showQrCode">Quasar v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>
    <q-dialog v-model="qrCodeShow">
      <q-card style="min-width: 20rem">
        <q-card-section>
          <div class="text-h6">二维码</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-select
            @update:model-value="ipChange"
            outlined
            v-model="ip"
            :options="options"
            label="ip地址"
          />
          <Qrcode :ip="ip"></Qrcode>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="OK" color="primary" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>
    <q-drawer elevated v-model="leftDrawerOpen" bordered>
      <q-list>
        <q-item-label header> Essential Links </q-item-label>

        <EssentialLink
          v-for="link in essentialLinks"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from "components/EssentialLink.vue";

const linksList = [
  {
    title: "Docs",
    caption: "quasar.dev",
    icon: "school",
    link: "https://quasar.dev",
  },
  {
    title: "Github",
    caption: "github.com/quasarframework",
    icon: "code",
    link: "https://github.com/quasarframework",
  },
  {
    title: "Discord Chat Channel",
    caption: "chat.quasar.dev",
    icon: "chat",
    link: "https://chat.quasar.dev",
  },
  {
    title: "Forum",
    caption: "forum.quasar.dev",
    icon: "record_voice_over",
    link: "https://forum.quasar.dev",
  },
  {
    title: "Twitter",
    caption: "@quasarframework",
    icon: "rss_feed",
    link: "https://twitter.quasar.dev",
  },
  {
    title: "Facebook",
    caption: "@QuasarFramework",
    icon: "public",
    link: "https://facebook.quasar.dev",
  },
  {
    title: "Quasar Awesome",
    caption: "Community Quasar projects",
    icon: "favorite",
    link: "https://awesome.quasar.dev",
  },
];

import { defineComponent, ref } from "vue";

import { onMounted } from "vue";
import { api } from "boot/axios";
import Qrcode from "components/Qrcode.vue";

export default defineComponent({
  name: "MainLayout",

  components: {
    Qrcode,
    EssentialLink,
  },

  setup() {
    const leftDrawerOpen = ref(false);
    let qrCodeShow = ref(false);
    let ip = ref("");
    let options = ref([]);

    function showQrCode() {
      ip.value = options.value[0];
      qrCodeShow.value = true;
    }

    function ipChange(value) {
      console.log(value);
      ip.value = value;
    }

    onMounted(() => {
      api.get("/ips").then(({ data }) => {
        console.log(data);
        options.value = data.ips;
      });
    });
    return {
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
      showQrCode,
      qrCodeShow,
      ip,
      options,
      ipChange,
    };
  },
});
</script>
