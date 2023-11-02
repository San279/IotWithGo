<template>
    <div class = "remoteContainer">
        <div class = "remote">
            <div class = "screen">
            </div>
            <div class = "buttons">
                <img class="power" src="../assets/power-on.png"/>
                <br/>
                <br/>
                <button class = "cool"> COOL</button>
                <button class = "dry" > DRY </button>
                <br/>
                <br/>
                <button class = "increase"> + </button>
                <br/>
                <br/>
                <img class="fan" src="../assets/fan.png"/>
                <br/>
                <br/>
                <button class = "decrease"> - </button>
                <br/>
                <br/>
                <input class = "timer" type = "number" placeholder = "   Timer/mins" />
            </div>
        </div>
    </div>
</template>

<script>
    export default{
        name: "HaierOne",
        data(){
            return{
                id: 0,
                username:"",
                irId:"",
                air: "",
                fan: "",
                temp: 0,
                ir: [],
                timer: 0,
        }
        }
    }
</script>

<style>

.remoteContainer {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 900px;
    display: flex;
    background-color: white;
}

.remote {
    width: 300px;
    border-radius: 10px;
    height: 700px;
    display: inline;
    border: 0.5px solid black;
    background-color: white;
}

.screen {
    width: 90%;
    height: 37%;
    margin: auto;
    border: 1px solid white;
    margin-top: 15px;
    background-color: black;
}

.power {
    padding-left: 37.5%;
    padding-top: 10%;
    width: 25%; 
    height:10%;
}

.power:hover{
    opacity: 0.3;
}

.cool{
    width: 30%;
    height: 30px;
    margin-right: 20%;
    margin-left: 10%;
    font-size: 15px;
    letter-spacing: 1.5px;
    background-color: lightgray;
    border: none;
    border-radius: 5px;
}

.dry{
    width: 30%;
    height: 30px;
    font-size: 15px;
    letter-spacing: 1.5px;
    background-color: lightgray;
    border: none;
    border-radius: 5px;
}

.increase{
    width: 15%;
    margin-left: 41%;
    font-size: 25px;
    background-color: lightgray;
    border: none;
}

.fan{
    width: 25%;
    opacity: 0.8;
    margin-left: 37%;
}

.fan:hover{
    opacity: 0.3; 
}


.decrease{
    width: 15%;
    margin-left: 41%;
    font-size: 25px;
    background-color: lightgray;
    border: none;
}

.timer{
    width: 30%;
    height: 15px;
    text-align: center;
    margin-left: 30.5%;
}



</style>