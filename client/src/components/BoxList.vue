<script setup lang="ts">
const props = defineProps<{ list: Set<string> }>()
</script>
<template>
  <ul class="ingredients">
    <li v-for="listElem in props.list" :key="listElem">
      {{ listElem }}
    </li>
  </ul>
</template>
<style scoped>
.ingredients {
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  padding: 0;
  margin: 0;
  justify-content: space-evenly;
  list-style: none;
  gap: 0.25rem;
}

.ingredients li {
  background: var(--highlight);
  border-radius: var(--border-radius);
  padding: 0.5rem 1rem;
  cursor: default;
}
</style>
