<script setup lang="ts">
import { recipesStore } from "../../stores/recipes.store";
import axios from "axios";
import { reactive } from "vue";

const excludes = reactive<Set<string>>(new Set());

const toggleExclude = (key: string) => {
  if (excludes.has(key)) {
    excludes.delete(key);
  } else {
    excludes.add(key);
  }
};

const sendRecipes = () => {
  const groceryList = Object.entries(recipesStore.ingredients).reduce<string[]>(
    (acc, [key, value]) => {
      if (excludes.has(key)) {
        return acc;
      }
      const amount = value.reduce<string>((localAcc, details) => {
        if (localAcc !== "") {
          localAcc += ", ";
        }
        localAcc += `${details.amount} ${details.unit}`;
        return localAcc;
      }, "");
      acc.push(`${key} - ${amount}`);
      return acc;
    },
    []
  );
  axios
    .post(`${import.meta.env.VITE_BACKEND_URL}/recipe/send`, {
      list: groceryList,
    })
    .catch((err) => console.error(err));
};
</script>
<template>
  <h1>Saved recipes</h1>
  <div class="recipes">
    <div
      v-for="recipe in recipesStore.recipes"
      :key="recipe.slug"
      class="recipe"
    >
      <span>{{ recipe.title }}</span>
      <button @click="recipesStore.removeRecipe(recipe)">x</button>
    </div>
  </div>
  <h2>Shopping list</h2>
  <ul class="list">
    <li
      v-for="entry in Object.entries(recipesStore.ingredients).sort()"
      :class="excludes.has(entry[0]) && 'exclude'"
      :key="entry[0]"
    >
      <button class="list-button" @click="toggleExclude(entry[0])">
        <strong>{{ entry[0] }}</strong> -
        <template v-for="ingredient in entry[1]"
          >{{ ingredient.amount }} {{ ingredient.unit }}</template
        >
      </button>
    </li>
  </ul>
  <button
    v-if="Object.entries(recipesStore.ingredients).length"
    @click="sendRecipes()"
  >
    Send grocery list to Slack!
  </button>
  <div v-else>No recipes added.</div>
</template>
<style scoped>
.recipes {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.exclude .list-button {
  text-decoration: line-through;
}

.recipe {
  padding: 0.5rem 1rem;
  background: #f39f8644;
  border-radius: var(--border-radius);
  text-align: left;
  gap: 1rem;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.list {
  text-align: left;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding: 0;
}

.list li {
  position: relative;
  list-style: none;
  padding-left: 2rem;
}

.list-button {
  flex: 1;
  background: none;
  padding: 0.5rem;
  border: none;
  outline: none;
  text-align: left;
  font-size: inherit;
  cursor: pointer;
  position: relative;
}

.list li:nth-child(odd) {
  background: #fff1;
}

.list-button:before {
  content: "o";
  position: absolute;
  top: 0.5rem;
  left: -1rem;
}

.exclude .list-button:before {
  content: "✔";
  color: green;
}
</style>
