<script setup lang="ts">
import { recipesStore } from "../../stores/recipes.store";
import axios from "axios";
import RecipeIngredients from "./RecipeIngredients.vue";
import IconElement from '../Icons/IconElement.vue'

const sendRecipes = () => {
  const groceryList = Object.entries(recipesStore.ingredients).reduce<string[]>(
    (acc, [key, value]) => {
      if (recipesStore.excludes.has(key)) {
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
      <button class="remove" @click="recipesStore.removeRecipe(recipe)" :title="`Remove ${recipe.title} from saved list`">
        <icon-element name="minus"/>
        <span>Remove from list</span>
      </button>
    </div>
  </div>
  <h2>Shopping list</h2>
  <p>Select which items should be on the grocery list</p>
  <recipe-ingredients></recipe-ingredients>
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

.remove {
  font-size: 0.75rem;
  gap: 0.5rem;
}

.recipe {
  padding: 0.5rem 1rem;
  background: var(--background-brown);
  border-radius: var(--border-radius);
  text-align: left;
  gap: 1rem;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
</style>
