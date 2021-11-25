<template>
  <div id="qrcode">

  </div>
</template>

<script>

export default {
  name: "Qrcode",
  props:{
    ip:{
      type:String
    }
  },data(){
    return {
      qrcode:''
    }
  },
  watch:{
    ip  (val){
      if (val ) {

        this.qrcode.clear();
        this.qrcode.makeCode(val);
      }
    }
  },
  mounted() {
    this.qrcode=new QRCode(document.getElementById("qrcode"), {
      width: 200,
      height: 200
    });
    this.qrcode.clear();
    this.qrcode.makeCode(this.ip);
  }
}
</script>

<style scoped>

</style>
