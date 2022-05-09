import { boot } from "quasar/wrappers";
import dayjs from "dayjs";
import VConsole from "vconsole";

export default boot(({ app, router }) => {
  // something to do

  const vConsole = new VConsole();
  app.config.globalProperties.$dayjs = dayjs;
});
