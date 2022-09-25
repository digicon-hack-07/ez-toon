<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div :class="$style.thumbnail">
    <div :class="$style.area">
      <img src="/vite.svg" :class="$style.image" />
      <p
        ref="nameRef"
        :class="[$style.name, isHidden ? $style.overflowName : '']"
      >
        {{ name }}
      </p>
      <span :class="[isHidden ? $style.overflow : $style.hide]">...</span>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'

const isHidden = ref<boolean>(false)

withDefaults(
  defineProps<{
    name: string
    image: string
  }>(),
  {
    name: 'noname',
    image: '/vite.svg'
  }
)

const nameRef = ref<HTMLElement>()
onMounted(() => {
  const fontSize = getComputedStyle(document.documentElement).fontSize
  const remHeight = (nameRef.value?.clientHeight ?? 0) / parseFloat(fontSize)
  isHidden.value = remHeight > 3
  //alert(remHeight)
})
</script>

<style module>
.image {
  width: 8rem;
  height: 8rem;
}

.thumbnail {
  display: inline-block;
  vertical-align: top;
  text-align: center;
  border-radius: 0.5rem;
  padding: 1rem;
  margin: 2rem;
}

.thumbnail:hover {
  background-color: darkgray;
}

.name {
  line-height: 1.3rem;
  width: 8rem;
}

.overflowName {
  overflow: hidden;
  text-overflow: ellipsis;
  overflow-wrap: normal;
  height: 2.6rem;
}

.overflow {
  background: #fff;
  font-size: 1rem;
  width: 1rem;
  height: 1rem;
  /* position: absolute; */
  position: relative;
  left: 3.5rem;
  top: -2.35rem;
  padding: -0.5rem;
  margin-top: 1.3rem;
}

.thumbnail:hover .overflow {
  background-color: darkgray;
}

.area {
  height: 12rem;
}

.hide {
  display: none;
}
</style>
