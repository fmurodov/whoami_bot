import logging
import os
from dotenv import load_dotenv
from aiogram import Bot, Dispatcher, executor, types

load_dotenv()
API_TOKEN = os.getenv("API_TOKEN")

# Configure logging
logging.basicConfig(level=logging.INFO)

# Initialize bot and dispatcher
bot = Bot(token=API_TOKEN)
dp = Dispatcher(bot)


@dp.message_handler(commands=['start', 'help'])
async def send_welcome(message: types.Message):
    """
    This handler will be called when user sends `/start` or `/help` command
    """
    await message.reply(f"Hi!\nI'm whoamiBot!\nchat id: {message.chat.id}",)


@dp.message_handler(commands=['whoami', 'me'])
async def send_welcome(message: types.Message):
    """
    This handler will be called when user sends `/whoami` or `/me` command
    """
    await message.reply(f'Your data: \n'
                        f'chat id: {message.chat.id}\n'
                        f'username: {message.chat.username}\n'
                        f'name: {message.chat.first_name} {message.chat.last_name}',)


@dp.message_handler()
async def echo(message: types.Message):
    await message.answer(message.text)


if __name__ == '__main__':
    executor.start_polling(dp, skip_updates=True)
