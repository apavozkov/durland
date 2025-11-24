# Durland

## World definition
- Локации, фауна и эффекты окружения описаны в `world_definition.json`. Добавление новой локации или местности сводится к дописанию блока в этот файл.
- `models.LoadWorldDefinition(path)` читает JSON; ожидается, что файл указан и не пустой.
- Конвертация описания в состояние мира: `def.ToWorldState()`.

## Races
- Расы и народы описаны в `races.json`.

## Activities
- Активности описаны в `activities.json`.

## Пояснение

### Все эффекты описываются в JSON и применяются через систему регистров:

Регистр обработчиков эффектов: `effectRegistry map[string]EffectHandler`

Регистр проверки условий: `conditionRegistry map[string]ConditionChecker`

### Вся логика вынесена в конфигурационные файлы:

Расы и народы с их уникальными эффектами: `races.json` 

Активности и их базовые эффекты: `activities.json`

Локации, местности и их особенности: `world_definition.json` 

## Легкое дополнение

### Для добавления новых "сущностей" достаточно описать их в JSON

Добавим новую расу и народ. Пример:

Добавим в JSON:
```
{
  "name": "Маги",
  "peoples": [
    {
      "name": "Волшебники",
      "effects": [
        {
          "type": "multiply_all_changes",
          "conditions": [{"activity_is": "zumbalit"}],
          "parameters": {"multiplier": 1.5}
        }
      ]
    }
  ]
}
```
### *Для добавления нового эффекта нужно добавить обработчик в регистр

Добавим новый эффект вероятной потери денег. Пример:

Добавим обработчик в регистр:
```
ec.effectRegistry["steal_money"] = func(result *EffectResult, durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) {
    if probability, ok := params["probability"].(float64); ok && rand.Float64() < probability {
        if amount, ok := params["amount"].(float64); ok {
            result.MoneyChange -= amount
        }
    }
}
```

После добавления обработчика можем обновить JSON:
```
{
  "type": "steal_money",
  "parameters": {
    "probability": 0.2,
    "amount": 5.0
  }
}
```

### *Для добавления нового условия нужно добавить проверку в регистр

Мы хотим восполнять здоровье при достижении определенного порога. Пример:

Добавим проверку уровня здоровья:
```
ec.conditionRegistry["has_low_health"] = func(durlian *Durlian, activity string, worldState *WorldState, params map[string]interface{}) bool {
    if threshold, ok := params["threshold"].(float64); ok {
        return durlian.Stats.Health <= threshold
    }
    return false
}
```

Добавим в JSON:
```
{
  "type": "add_health",
  "conditions": [
    {
      "type": "has_low_health",
      "parameters": {"threshold": 3.0}
    }
  ],
  "parameters": {"value": 5.0}
}
```
## Доступные эффекты

Прибавляет значение к показателю:
- add_health, add_money, add_satisfaction

Умножает изменение показателя:
- multiply_health_change, multiply_money_change, multiply_satisfaction_change

Умножает все изменения:
- multiply_all_changes

Эффекты, зависящие от количества существ:
- fauna_based_health, fauna_based_money, fauna_based_satisfaction
- chance_fauna_money_loss, chance_fauna_damage

Вероятностные эффекты:
- chance_health_save 
- chance_money_wipe

Удовлетворенность, зависящая от фауны в истории перемещений:
- history_based_satisfaction

Устанавливает конкретное значение изменения (например, обнулить)
- set_stat_change

## Доступные условия
Эффект применяется только для указанной активности:
- activity_is

Эффект только для указанного народа:
- people_is

Эффект только в указанной локации:
- location_is

Эффект срабатывает после N шагов в локации:
- min_stay_count
