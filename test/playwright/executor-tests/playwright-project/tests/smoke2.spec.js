// @ts-check
const { test, expect } = require("@playwright/test");

test("Smoke 2 - has title", async ({ page }) => {
  await page.goto("https://beneva.ca/");

  await expect(page).toHaveTitle(/Beneva | Insurance & Financial Services/);
});
