import { boot } from 'quasar/wrappers'
import dayjs from 'dayjs'
export default boot(  ({ app, router}) => {
  // something to do
  app.config.globalProperties.$dayjs = dayjs;
})
