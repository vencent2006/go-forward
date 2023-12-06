export default {
  namespaced: true,
  // () => ({sidebarOpened: true}), 这么写才能返回一个对象
  // () => {}, 这么写是单纯的执行一个函数，并没有返回值
  state: () => ({
    sidebarOpened: true
  }),
  mutations: {
    triggerSidebarOpened(state) {
      state.sidebarOpened = !state.sidebarOpened
    }
  },
  actions: {}
}
