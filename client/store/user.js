import { createStore } from "vuex";

export default createStore({
  state() {
    return {
      count: 0,
      user: "",
      token: "",
      isLogin: false,
    }
  },
  mutations: {
    setUser(state, u) {
      state.user = u;
      state.isLogin = true;
    },
    setToken(state, t) {
      state.token = t;
    },
    logout(state){
      state.user = "";
      state.token = "";
      state.isLogin = false;
    }
  },

  actions: {
    setUser ({ commit }, u ){
        commit('setUser', u);
    },
    setToken ({ commit },t) {
        commit('setToken', t);
    },

  },

})
/*
export default createStore({
  state,
  actions,
  mutations
})
*/
