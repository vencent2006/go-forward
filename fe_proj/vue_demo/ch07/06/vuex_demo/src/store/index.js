import { createStore } from 'vuex'
// VueX 数据管理框架
// VueX 创建了一个全局唯一的仓库，用来存放全局的数据
export default createStore({
  state: {
    name: 'dell'
  },
  getters: {
  },
  // mutation 里面只允许写同步代码，不允许写异步代码
  // commit 和 mutation 做关联
  mutations: {
    // 第四步，对应的mutation 被执行
    change(state, str) {
      console.log('mutation change')
      // 第五步，在mutation 里面修改数据
      state.name = str
    }
  },
  // action 可以放异步的代码
  // dispatch 和 action 做关联
  actions: {
    // 第二步， store 感知到你触发了一个叫做 change 的action，执行change
    change(store, str) {
      // console.log(store)
      console.log('actions change')
      // 第三步，提交一个 commit 触发一个mutation
      setTimeout(() => {
        this.commit('change', str)
      }, 2000);
    }
  },
  modules: {
  }
})
