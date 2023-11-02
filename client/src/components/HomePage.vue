<script setup>
    function OpenHaier1(){
        window.open("http://192.168.1.30:8080/ac1", "_self");
    }

</script>




<template>
    <div class = "imgContainer">
        <img class="acRemote" src="../assets/59136.jpg" @click = "OpenHaier1()"/>
        <img class="waterPump" src="../assets/watering-plants_7799888.png" />
    </div>
</template>

<style>
    .imgContainer {
        padding: 80px;
        display: flex;
        justify-content: space-evenly;
    }

    .imgContainer img{
        width: 30%;
    }


    .acRemote:hover{ 
      opacity: 0.7;
    }

    .waterPump{
        opacity: 0.4;
    }


</style>
