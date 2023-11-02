<template>
    <div class = "loginPage">
        <img class="logo" src="../assets/logo.png" />
        <br/>
        <br/>
        <div class = "login">
            <input type = "text" v-model = "username" placeholder = "Username or Mail" />
            <br/>
            <br/>
            <input type = "password" v-model = "password" placeholder = "Password" />
            <br/>
            <br/>
            <br/>
            <button v-on:click = "login"> Login</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
//import { useStore } from 'vuex';
//import { computed } from 'vue'
export default {
        
        name :'LoginPage',
        data(){   
            return {
                id: 0,
                username: "",
                password: ""
            }
    },
    methods:{
            async login(){  

                const RequestJson = JSON.stringify({
                    id: 0,
                    username : this.$data.username,
                    password : this.$data.password
                })
                try{
                    let result = await axios.post("http://localhost:5001/login", RequestJson, {
                        headers: {
                            'Content-Type':'application/json; charset=UTF-8',
                        }
                    });
                    console.log(result);
                    if(result.status === 200){
                        console.log("login successfully");
                        this.$store.dispatch('setUser',this.$data.username);
                        this.$store.dispatch('setToken', result.data.token);
                    }
                }catch(err){
                    console.log(err);
                }
                console.log(this.$store.state.user);
                console.log(this.$store.state.token);
                console.log(this.$store.state.isLogin);
            },
     },
};

</script>

<style>
.loginPage{
    width: 100%;
    height: 800x;
    margin-top: 40px;
    text-align: center;
}
.logo{
    width: 300px;
    margin-left: 0 auto;
    margin-right: auto;
}
.login {
    align-items: center;
    align-content: center;
}
.login input {
    width: 300px;
    height: 20px;
    align-items: center;
    border: 1px solid;
    padding: 10px;
    border-radius: 5px;
    border-color: #0080ff;
}

.login button {
    width: 200px;
    height: 40px;
    border: none;
    color: white;
    background-color: #0080ff;
    border-radius: 5px;
    
}
</style>