import { reactive } from 'vue'
import { Ingredient } from '../models'

const defaultIngredients = [{ name: '', amount: null, unit: '' }]

interface RecipeStore {
  ingredients: Ingredient<number | null>[]
  addIngredient: () => void
  setIngredients: (ingredients?: Ingredient[]) => void
  removeIngredient: (index: number) => void
  clear: () => void
}

export const recipeStore = reactive<RecipeStore>({
  ingredients: [{ name: '', amount: null, unit: '' }],
  addIngredient() {
    this.ingredients.push({ name: '', amount: null, unit: '' })
  },
  setIngredients(ingredients?: Ingredient[]) {
    if (ingredients) {
      this.ingredients = ingredients
    }
  },
  removeIngredient(index) {
    this.ingredients.splice(index, 1)
  },
  clear() {
    this.ingredients = defaultIngredients
  },
})
