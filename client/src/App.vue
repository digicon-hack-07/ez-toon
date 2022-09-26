<template>
   <main>
    <div :class="$style.container">
      <header :class="$style.header">
        <router-link to="/">Menu</router-link>
      </header>
      <router-view />
    </div>
  </main>
  <!-- <hello-world msg="Vite + Vue" /> -->
</template>

<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
//import HelloWorld from './components/HelloWorld.vue'
</script>

<style module>
.container {
  margin: auto;
  text-align: center;
}
.header {
  height:2rem;
  margin-bottom: 1rem;
  display: flex;
  justify-content: center;
}
</style>
