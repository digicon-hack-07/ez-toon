<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div :class="$style.thumbnail">
    <div :class="$style.area">
      <div v-if="isImageExists">
        <img :src="image" :class="$style.image" />
      </div>
      <div v-else>
        <img :src="'/noImage.svg'" :class="$style.image" />
      </div>
      <p
        ref="nameRef"
        :class="[$style.name, isHidden ? $style.overflowName : '']"
      >
        {{ name }}
      </p>
      <span :class="[isHidden ? $style.overflow : $style.hide]">...</span>
      <p v-if="isShowDate" :class="$style.name">{{ dateString }}</p>
      <div v-if="isShowArrow" :class="$style.arrowSpace">
        <img
          src="/leftArrow.svg"
          :class="$style.leftArrow"
          @click.stop="$emit('left')"
        />
        <img
          src="/rightArrow.svg"
          :class="$style.rightArrow"
          @click.stop="$emit('right')"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, nextTick, onMounted, ref, watch } from 'vue'

const isHidden = ref<boolean>(false)

const props = withDefaults(
  defineProps<{
    name: string
    image: string
    isShowDate: boolean
    isShowArrow: boolean
    createdAt?: Date
    updatedAt?: Date
  }>(),
  {
    name: 'noname',
    image: '/noImage.svg',
    isShowDate: false,
    isShowArrow: false
  }
)
interface Emits {
  (e: 'left'): void
  (e: 'right'): void
}
const isImageExists = ref(true)

defineEmits<Emits>()

const nameRef = ref<HTMLElement>()

const formatDate = (dt: Date) => {
  var y = ('' + dt.getFullYear())
  var m = dt.getMonth() + 1
  var d = dt.getDate()
  var h = ('00' + dt.getHours()).slice(-2);
  var min = ('00' + dt.getMinutes()).slice(-2);
  return y + '/' + m + '/' + d + ' ' + h + ':' + min
}

onMounted(() => {
  const fontSize = getComputedStyle(document.documentElement).fontSize
  const remHeight = (nameRef.value?.clientHeight ?? 0) / parseFloat(fontSize)
  isHidden.value = remHeight > 3
  const img = new Image()
  img.src = props.image
  img.onerror = () => {
    isImageExists.value = false
  }

  //alert(remHeight)
})

watch(props, () => {
  isHidden.value = false
  nextTick(() => {
    const fontSize = getComputedStyle(document.documentElement).fontSize
    const remHeight = (nameRef.value?.clientHeight ?? 0) / parseFloat(fontSize)
    isHidden.value = remHeight > 3
  })
})

const dateString = computed(() => {
  const displayDate = props.createdAt ?? new Date(10000000000)
  return formatDate(displayDate)
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
    background-color:#f0f0f0;
}

.name {
  line-height: 1.3rem;
  width: 8rem;
  margin: 0;
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
  top: -1.4rem;
  padding: -0.5rem;
  margin-top: 1.3rem;
}

.thumbnail:hover .overflow {
    background-color:#f0f0f0;
}

.area {
  height: 12rem;
}

.hide {
  display: none;
}

.arrowSpace {
  width: 100%;
  height: 3rem;
}
.leftArrow {
  height: 3rem;
  margin-right: 1rem;
}

.rightArrow {
  height: 3rem;
  margin-left: 1rem;
}

.leftArrow:hover{
  filter: drop-shadow(0 0 0.5em #606060);
  border-radius: 0.3rem;
}

.rightArrow:hover{
  filter: drop-shadow(0 0 0.5em #606060);
  border-radius: 0.3rem;
}
</style>
