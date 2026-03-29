"""Minimal Aiogram probe; production workload lives in cmd/bot (Go)."""
import asyncio
import os

from aiogram import Bot, Dispatcher, types

TOKEN = os.environ.get("TELEGRAM_BOT_TOKEN", "")

bot = Bot(token=TOKEN)
dp = Dispatcher()


@dp.message()
async def handler(message: types.Message):
    await message.answer("Bot is up (Python probe). Use the Go service for full features.")


async def main():
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())
