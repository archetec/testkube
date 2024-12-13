// @ts-check
const { test, expect } = require("@playwright/test");

test("Smoke 1 - has title", async ({ page }) => {
  await page.goto("https://www.beneva.ca/en");

  await expect(page).toHaveTitle(/Beneva | Insurance & Financial Services/);
});
