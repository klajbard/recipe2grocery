<script setup lang="ts">
import axios from "axios";
import { onMounted, reactive, ref } from "vue";
import { ApiResponse, ExtendedRecipe, Recipe } from "../../models";
import DetailsModal from "../DetailsModal.vue";
import RecipeItem from "./RecipeItem.vue";
let pageNum = 1;

const modalContentRecipe = ref<ExtendedRecipe | null>(null);
const recipes = reactive<ExtendedRecipe[]>([]);
const reachedEnd = ref(false);

const loadMore = async () => {
  axios
    .get<ApiResponse<Recipe[]>>(
      `${import.meta.env.VITE_BACKEND_URL}/recipes/${pageNum}`
    )
    .then((resp) => {
      if (resp.data == null || !resp.data.data) {
        throw new Error("No more data to fetch");
      }
      resp.data.data.forEach((recipe:Recipe) => {
        recipes.push({
          ...recipe,
          removed: false,
        })
      })
      pageNum += 1;
    })
    .catch((err) => {
      console.error(err);
      if (err.message === "No more data to fetch") {
        reachedEnd.value = true;
      }
    });
};

onMounted(() => {
  loadMore();
});

const handleModalClose = () => {
  modalContentRecipe.value = null
}

const handleModalOpen = (recipe: ExtendedRecipe) => {
  modalContentRecipe.value = recipe
}

const handleRemove = () => {
  if (modalContentRecipe.value) {
    modalContentRecipe.value.removed = true
  }
}
</script>
<template>
  <h1>Recipes</h1>
  <p>Browse and save recipes which will then be used to create the grocery list from.</p>
  <div class="recipes">
    <recipe-item v-for="recipe in recipes" :key="recipe.slug" :recipe="recipe" @show-overlay="handleModalOpen(recipe)">
    </recipe-item>
    <button v-if="!reachedEnd" @click="loadMore">Load more</button>
    <div class="empty" v-if="reachedEnd">No more recipe to load...</div>
  </div>
  <details-modal :recipe="modalContentRecipe" @modal-close="handleModalClose" @recipe-remove="handleRemove"></details-modal>
</template>
<style scoped>

.recipes {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.empty {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.6rem 1rem;
  border-radius: 8px;
  background: var(--button-background);
}
</style>
