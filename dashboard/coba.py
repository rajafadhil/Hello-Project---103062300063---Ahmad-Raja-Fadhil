import streamlit as st
import pandas as pd
import matplotlib.pyplot as plt
import os

# Tentukan folder tempat file-file dataset disimpan
folder_path = "C:/Users/RAJA/dicoding/dashboard"

# List untuk menyimpan DataFrames dari setiap file
dfs = []

# Loop melalui setiap file dalam folder
for filename in os.listdir(folder_path):
    if filename.endswith(".csv"):
        filepath = os.path.join(folder_path, filename)
        # Baca file dataset ke dalam DataFrame
        df = pd.read_csv(filepath)
        dfs.append(df)

# Gabungkan semua DataFrames menjadi satu DataFrame besar berdasarkan kolom 'station'
combined_df = pd.concat(dfs, ignore_index=True)

# Sekarang Anda memiliki DataFrame gabungan yang dapat Anda gunakan untuk analisis lebih lanjut


# Definisikan fungsi untuk membuat visualisasi data
# Hitung rata-rata CO per tahun
average_co_per_year = combined_df.groupby(['year', 'station'])['CO'].mean().reset_index()

city_names = combined_df['station'].unique().tolist()

# Dropdown menu di sidebar untuk memilih nama kota
selected_city = st.sidebar.selectbox("Pilih Nama Kota:", city_names)

# Filter DataFrame berdasarkan nama kota yang dipilih
filtered_data = average_co_per_year[average_co_per_year['station'] == selected_city]

st.title(f"Rata-rata CO per Tahun Kota {selected_city}")

# Membuat plot berdasarkan data yang telah difilter
st.bar_chart(filtered_data.set_index('year')['CO'])

# Filter DataFrame untuk kota yang dipilih
selected_city_data = combined_df[combined_df['station'] == selected_city]

# Membuat grafik scatter untuk tingkat curah hujan vs konsentrasi CO
plt.figure(figsize=(14, 6))
plt.scatter(selected_city_data['RAIN'], selected_city_data['CO'], color='blue')
plt.title("Pengaruh Curah Hujan terhadap Konsentrasi CO")
plt.xlabel("Curah Hujan")
plt.ylabel("Konsentrasi CO")
plt.grid(True)

# Menambahkan grid yang sama dengan plot sebelumnya
plt.gca().set_axisbelow(True)

# Menampilkan plot menggunakan Streamlit
st.pyplot(plt)

st.markdown("&copy; 2024 rajafadhil")