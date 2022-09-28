<template>
  <header :class="$style.header">
    <router-link to="/">Menu</router-link>
  </header>
  <h1 :class="$style.notFound">Page Not Found</h1>
</template>

<style module>
.notFound {
  width: 90vw;
  height: 80vh;
  margin: 0;
}

.header {
  height: 2rem;
  margin-bottom: 1rem;
  display: flex;
  justify-content: center;
  margin-top: 1.5rem;
  font-size: 1.5rem;
}
</style>
