<script setup>


import {TurnOff, TurnOn} from "../../wailsjs/go/main/App";

function turnOn() {
  TurnOn()
}

function turnOff() {
  TurnOff()
}

</script>

<template>
  <main>
    <!--    <div class="input-box">-->
    <div>

      <div style="margin: 10px">
        <n-button ghost dashed color="#8a2be2" type="info"  size="large" @click="turnOn"
                  style="width: 150px;height: 80px;font-size: 40px"> 开机
        </n-button>

      </div>
      <div style="margin: 50px">
        <n-button ghost dashed color="#ff69b4" type="error"  size="large" @click="turnOff"
                  style="width: 150px;height: 80px;font-size: 40px"> 关机
        </n-button>

      </div>
    </div>

    <!--    </div>-->
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box {
  text-align: center;
}

.input {
  width: 200px;
  margin-right: 10px;
}
</style>
