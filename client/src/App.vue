<template>
  <main>
    <div :class="$style.viewer">
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
.viewer {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  max-width: 100%;
  max-height: 100%;
  margin: auto;
}
</style>
