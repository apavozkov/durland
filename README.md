# Durland

## World definition
- Локации, фауна и эффекты окружения описаны в `world_definition.json`. Добавление новой локации или местности сводится к дописанию блока в этот файл.
- `models.LoadWorldDefinition(path)` читает JSON; ожидается, что файл указан и не пустой.
- Конвертация описания в состояние мира: `def.ToWorldState()`.
