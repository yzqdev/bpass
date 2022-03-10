import { boot } from 'quasar/wrappers'
import axios, {AxiosInstance} from 'axios'

import qs from 'qs'
declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
  }
}
export const baseUrl='http://localhost:8901'
const api = axios.create({ baseURL: baseUrl })
api.interceptors.request.use((config) => {

  return config
})
api.interceptors.response.use(
  function (response) {
    // 对响应数据做点什么
    let { data } = response;
    return data;
  },
  function (error) {
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);
export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
})

export { api }
