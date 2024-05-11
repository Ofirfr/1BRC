import csv
import random

# Read city names from the file
city_names = []
with open("./data/worldcities.csv", newline='', encoding="utf-8") as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
        city_names.append(row['city'])

# Shuffle the city names to ensure randomness
random.shuffle(city_names)

# Generate 999 different stations with real city names
stations = city_names[:999]

# Function to generate a random temperature measurement with one fractional digit
def generate_measurement():
    return round(random.uniform(-99.9, 99.9), 1)

# Generate temperature measurements for each station
with open("./data/temperature_data.txt", "w", encoding="utf-8") as file:
  row_count = 1000000 # Million!
  for _ in range (row_count):
       station = random.choice(stations)
       measurement = generate_measurement()
       file.write(f"{station};{measurement}\n")