<script setup lang="ts">
import { Ingredient } from "../../models";
import { recipeStore } from "../../stores/recipe.store";
import IconElement from "../Icons/IconElement.vue";
const units = [
  "gramm",
  "dkg",
  "kg",
  "liter",
  "dl",
  "ml",
  "cup",
  "tablespoon",
  "teaspoon",
  "bunch",
  "pinch",
  "pcs",
];

defineProps<{ ingredient: Ingredient<number | null>, index: number }>()
</script>
<template>
  <div class="ingredient">
    <label :for="`ingredient-amount-${index}`">Amount</label>
    <label :for="`ingredient-unit-${index}`">Unit</label>
    <label :for="`ingredient-name-${index}`">Name</label>
    <input :id="`ingredient-amount-${index}`" v-model="ingredient.amount" type="number" name="ingredient-amount" />
    <select :id="`ingredient-unit-${index}`" v-model="ingredient.unit" name="ingredient-unit">
      <option v-for="unit in units" :value="unit" :key="unit">
        {{ unit }}
      </option>
    </select>
    <input :id="`ingredient-name-${index}`" v-model="ingredient.name" name="ingredient-name" />
  </div>
  <button v-if="recipeStore.ingredients.length > 1" type="button" class="close"
    @click="recipeStore.removeIngredient(index)">
    <icon-element name="remove" />
  </button>
</template>
<style scoped>
.ingredient {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  width: 100%
}

.close {
  padding: 0.5rem 0.75rem;
  background: var(--color-orange-fade);
}
</style>
