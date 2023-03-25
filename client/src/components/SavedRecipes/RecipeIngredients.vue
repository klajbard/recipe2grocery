<script setup lang="ts">
import { recipesStore } from "../../stores/recipes.store";
import IconElement from "../Icons/IconElement.vue";

const toggleExclude = (key: string) => {
  if (recipesStore.excludes.has(key)) {
    recipesStore.excludes.delete(key);
  } else {
    recipesStore.excludes.add(key);
  }
};
</script>
<template>
  <div role="list" class="list">
    <label role="listitem" class="list-elem" :class="recipesStore.excludes.has(entry[0]) && 'exclude'"
      v-for="entry in Object.entries(recipesStore.ingredients).sort()" :key="entry[0]">
      <icon-element :name="recipesStore.excludes.has(entry[0]) ? 'checkbox' : 'checkbox-checked'"></icon-element>
      <input :checked=!recipesStore.excludes.has(entry[0]) type="checkbox" @change="toggleExclude(entry[0])">
      <strong>{{ entry[0] }}</strong> -
      <template v-for="ingredient in entry[1]">{{ ingredient.amount }} {{ ingredient.unit }}</template>
    </label>
  </div>
</template>
<style scoped>
.list {
  text-align: left;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding: 0;
}

.list-elem:nth-child(odd) {
  background: var(--highlight);
}

.list-elem {
  position: relative;
  cursor: pointer;
  user-select: none;
  display: flex;
  gap: 0.5rem;
  padding: 0.5rem;
  padding-left: 1rem;
  --icon-size: 1.5rem;
}

.list-elem:focus-within {
  outline: 2px solid #5E9ED6;
  border-radius: 3px;
}

.exclude {
  color: #fff7;
}

.list-elem input {
  position: absolute;
  height: 1rem;
  width: 1rem;
  left: 0.5rem;
  top: 0.5rem;
  opacity: 0;
}
</style>