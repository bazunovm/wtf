#!/usr/bin/env bash
set -e

# Проверка прав суперпользователя
if [ "$EUID" -ne 0 ]; then
  echo "Пожалуйста, запустите скрипт с sudo или от root"
  exit 1
fi

# Функция для проверки и установки golang
install_go() {
  if command -v go >/dev/null 2>&1; then
    echo "Go уже установлен"
    return
  fi

  echo "Go не найден. Пытаюсь установить..."

  if command -v apt >/dev/null 2>&1; then
    apt update
    apt install -y golang
  elif command -v dnf >/dev/null 2>&1; then
    dnf install -y golang
  elif command -v yum >/dev/null 2>&1; then
    yum install -y golang
  elif command -v pacman >/dev/null 2>&1; then
    pacman -Sy --noconfirm go
  else
    echo "Не удалось определить пакетный менеджер. Установите Go вручную."
    exit 1
  fi
}

# Установка Go
install_go


# Установка зависимостей и сборка
echo "Устанавливаю зависимости..."
go mod tidy

echo "Собираю бинарник..."
go build -o wtf ./cmd/wtf

# Создание папки для правил
echo "Создаю папку для правил..."
mkdir -p /usr/local/share/wtf

# Копирование rules.yaml
echo "Копирую rules.yaml..."
cp ./internal/explainer/rules.yaml /usr/local/share/wtf

# Копирование бинарника
echo "Устанавливаю бинарник..."
cp wtf /usr/local/bin/

echo "Установка завершена! Теперь можно запускать 'wtf' из любой директории."
