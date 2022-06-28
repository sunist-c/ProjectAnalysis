import time
import sys
import requests
import urllib3
import pandas as pd

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


def request(url):
    try:
        page = requests.get(url, verify=False, headers={
            "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"})
        if page.status_code == 200:
            return page
        else:
            print(page.status_code)
            return None
    except Exception as e:
        print(e)
        return None


def get_page(url):
    index = 1
    while True:
        page = request(url)
        if page is not None:
            with open(file='data.csv', mode='w') as w:
                w.write(page.text)
            break
        print("%s失败%d次" % (url, index))
        index += 1
        time.sleep(0.1)


def read_csv(date):
    data = pd.read_csv(filepath_or_buffer='data.csv', na_values={'Confirmed': 0, 'Deaths': 0, 'Recovered': 0})
    data.rename(
        columns={'Province/State': 'location_province', 'Province_State': 'location_province',
                 'Country/Region': 'location_country', 'Country_Region': 'location_country',
                 'Last_Update': 'refresh_time', 'Last Update': 'refresh_time', 'Confirmed': 'daily_confirm',
                 'Deaths': 'daily_death', 'Recovered': 'daily_recovered'},
        inplace=True)
    data.loc[:, 'refresh_time'] = date
    data.loc[data.location_country == 'US', 'location_country'] = 'United States'
    data.loc[data.location_country == 'Taiwan*', 'location_province'] = 'Taiwan'
    data.loc[data.location_country == 'Taiwan*', 'location_country'] = 'China'
    data.loc[data.location_country == 'Mainland China', 'location_country'] = 'China'
    for columns_name in ['daily_confirm', 'daily_death', 'daily_recovered']:
        data[columns_name].fillna(0, inplace=True)
    data.dropna(inplace=True, subset=['location_province', 'location_country', 'refresh_time'])
    data.drop(data[data['location_province'] == 'Unknown'].index, inplace=True)
    data.drop(data[data['location_country'] == 'Unknown'].index, inplace=True)
    data.to_csv(path_or_buf='data_result.csv', index=False, sep=',',
                columns=['location_country', 'location_province', 'refresh_time', 'daily_confirm', 'daily_death',
                         'daily_recovered'], header=False)


def main():
    if len(sys.argv) >= 2:
        date = sys.argv[1]
    else:
        date = '01-25-2020'
    get_page(
        'https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/{}.csv'.format(
            date))
    temp = date.split('-')
    date = temp[2] + '-' + temp[0] + '-' + temp[1]
    read_csv(date)


if __name__ == '__main__':
    main()
