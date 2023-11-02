<script setup>
  import { useStore } from 'vuex';
  const store = useStore();
  function getUser(){
    return store.state.user.toString();
  }
  function checkUser(){
    return store.state.isLogin;
  }
  function logout(){
    store.commit('logout');
  }
  
</script>

<template>
  <div class="topnav">
    <br/>
    <a><router-link to = "/home">HOME</router-link></a>
    <a> <router-link to = "/about">ABOUT</router-link> </a>

    <span v-if = "checkUser() === true"> 
         {{getUser()}}
    </span>
    <button v-if = "checkUser() === true" v-on:click = logout()>LOGOUT</button>
    <button v-if = "checkUser() === false">
      <router-link to = "/login">LOGIN</router-link>
    </button>
  </div>

  <router-view></router-view>
</template>

<script>
    export default {
    data() {
        return {
          activeComp: ''
        }
      }
    }
</script>

<style>
  #app {
      font-family: Avenir, Helvetica, Arial, sans-serif;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
  }

    .topnav {
      overflow: hidden;
      height: 60px;
      width: 100.8%;
      text-align: center;
      margin-left: -8px;
      background-color: #0080ff;
      margin-top: -8px;
    }

    /* Style the links inside the navigation bar */
    .topnav a {
      float: center;
      color: white;
      font-weight: 500;
      letter-spacing: 2px;
      padding: 20px;
      text-decoration: none;
      font-size: 15px;
    }

    .topnav span {
      float: center;
      color: white;
      font-weight: 500;
      position: absolute;
      margin-left: 31%;
      letter-spacing: 2px;
      text-decoration: none;
      font-size: 15px;
    }

    .topnav button {
      float: center;
      color: white;
      font-weight: 500;
      position: absolute;
      background-color: transparent;
      border: none;
      margin-left: 37%;
      letter-spacing: 2px;
      text-decoration: none;
      font-size: 15px;
    }



    /* Change the color of links on hover */
    .topnav a:hover {
      color: darkslategrey;
    }

    /* Add a color to the active/current link */
    .topnav a.active {
      background-color: #04AA6D;
      color: white;
      padding: 14px 16px;
    }

</style>
