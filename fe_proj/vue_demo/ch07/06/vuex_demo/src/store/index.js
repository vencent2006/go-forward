import { createStore } from 'vuex'
// VueX 数据管理框架
// VueX 创建了一个全局唯一的仓库，用来存放全局的数据
export default createStore({
  state: { name: 'dell' },
  mutations: {
    changeName(state, str) {
      console.log(str)
      state.name = str
    }
  },
  actions: {
    getData(store) {
      setTimeout(() => {
        store.commit('changeName', 'hello')
      }, 2000);
    },
  }
})
