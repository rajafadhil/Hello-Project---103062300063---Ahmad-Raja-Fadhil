import streamlit as st
import pandas as pd
import matplotlib.pyplot as plt
import os


# Tentukan folder tempat file-file dataset disimpan
folder_path = "dashboard/data_dashboard"

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
# Combine dataframes
combined_df = pd.concat(dfs)
combined_df.reset_index(drop=True, inplace=True)
# Filter DataFrame untuk tahun 2017
df_2017 = combined_df.loc[combined_df['year'] == 2017]

# Mengelompokkan data berdasarkan tahun dan stasiun, kemudian menghitung rata-rata CO
average_co_per_year = df_2017.groupby(['year', 'station'])['CO'].mean().reset_index()

# Dropdown menu di sidebar untuk memilih nama kota
selected_city = st.sidebar.selectbox("Pilih Nama Kota:", average_co_per_year['station'].unique())

# Filter DataFrame berdasarkan nama kota yang dipilih
filtered_data = average_co_per_year[average_co_per_year['station'] == selected_city]

# Menampilkan judul yang sesuai dengan kota yang dipilih
st.title(f"Rata-rata CO Tahun 2017 Kota {selected_city}")

# Membuat plot berdasarkan data yang telah difilter
st.bar_chart(filtered_data.set_index('year')['CO'])

# Membuat plot untuk rata-rata CO pada tahun 2017 untuk setiap kota
average_co_per_year_all_cities = combined_df.groupby(['year', 'station'])['CO'].mean().unstack()
average_co_2017 = average_co_per_year_all_cities.loc[2017]
st.title('Rata-rata CO Tahun 2017 pada Berbagai Kota')
st.bar_chart(average_co_2017)

# Filter DataFrame untuk kota yang dipilih
selected_city_data = combined_df[combined_df['station'] == selected_city]

# Membuat grafik scatter untuk tingkat curah hujan vs konsentrasi CO
st.title("Pengaruh Curah Hujan terhadap Konsentrasi CO")
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
