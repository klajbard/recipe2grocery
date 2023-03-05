<script setup lang="ts">
import axios from "axios";
import { ref, watch } from "vue";
import { ApiResponse, Recipe } from "../../models";
import { recipesStore } from "../../stores/recipes.store";
const props = defineProps<{ recipe: Recipe }>();

const isAdded = ref<boolean>(false);
const isRemoved = ref<boolean>(false);
const container = ref<HTMLDivElement | null>(null);
watch(recipesStore.recipes, () => {
  isAdded.value = !!recipesStore.recipes.find(
    ({ slug }) => slug === props.recipe.slug
  );
});

const handleUndo = () => {
  axios
    .post<ApiResponse>(`${import.meta.env.VITE_BACKEND_URL}/recipe`, {
      title: props.recipe.title,
      description: props.recipe.description,
      ingredients: props.recipe.ingredients,
    })
    .then((resp) => {
      if (resp.data.success) {
        isRemoved.value = false;
      }
    });
};

const handleRemove = () => {
  if (confirm(`Are you sure you want to remove "${props.recipe.title}"?`)) {
    axios
      .delete<ApiResponse>(
        `${import.meta.env.VITE_BACKEND_URL}/recipe/${props.recipe.id}`
      )
      .then((resp) => {
        if (resp.data.success) {
          isRemoved.value = true;
        }
      });
  }
};

const showOverlay = () => {
  console.log("Show more details of recipe");
};
</script>
<template>
  <div v-if="isRemoved">
    <button class="undo" @click="handleUndo()">
      Undo removal of {{ recipe.title }}
    </button>
  </div>
  <div v-else class="recipe" ref="container">
    <div class="content">
      <h3 class="title">{{ recipe.title }}</h3>
      <div class="box">
        <ul class="ingredients">
          <li v-for="ingredient in recipe.ingredients" :key="ingredient.name">
            {{ ingredient.name }}
          </li>
        </ul>
      </div>
    </div>
    <div class="button-group">
      <button @click="showOverlay()">Read more</button>
      <button @click="handleRemove()">Remove</button>
      <button
        class="add-remove-button"
        @click="
          isAdded
            ? recipesStore.removeRecipe(recipe)
            : recipesStore.addRecipe(recipe)
        "
      >
        {{ isAdded ? "-" : "+" }}
      </button>
    </div>
  </div>
</template>
<style scoped>
.recipe {
  padding: 0.5rem 1rem;
  background: #f39f8644;
  border-radius: var(--border-radius);
  text-align: left;
  gap: 1rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
}

.title {
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
}

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
  background: #fff1;
  border-radius: var(--border-radius);
  padding: 0.5rem 1rem;
  cursor: default;
}

.add-remove-button {
  width: 4rem;
}
.undo {
  width: 100%;
}
</style>
