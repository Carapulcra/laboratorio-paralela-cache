#include <iostream>
#include <chrono>
#include <cstdlib>
using namespace std;
using namespace std::chrono;

const int MAX = 10000;

double A[MAX][MAX], x[MAX], y[MAX];

void iniciar() {
    srand(42);
    for(int i = 0; i < MAX; i++){
        for(int j = 0; j < MAX; j++) A[i][j] = rand() % 100;
        x[i] = i;
        y[i] = 0;
    }
}

void first_pair() {
    for(int i = 0; i < MAX; i++){
        for(int j = 0; j < MAX; j++) y[i] += A[i][j] * x[j];
    }
}

void second_pair() {
    for(int j = 0; j < MAX; j++){
        for(int i = 0; i < MAX; i++) y[i] += A[i][j] * x[j];
    }
}

int main() {
    iniciar();
    auto start = high_resolution_clock::now();
    first_pair();
    auto stop = high_resolution_clock::now();
    auto duration = duration_cast<milliseconds>(stop - start);
    cout << "Tiempo first pair: " << duration.count() << "ms" << endl;
    iniciar();
    start = high_resolution_clock::now();
    second_pair();
    stop = high_resolution_clock::now();
    duration = duration_cast<milliseconds>(stop - start);
    cout << "Tiempo second pair: " << duration.count() << "ms" << endl;
}