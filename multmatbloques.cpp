#include <iostream>
#include <chrono>
#include <cstdlib>
using namespace std;
using namespace std::chrono;

const int MAX = 2000; //Va a variar
const int BLOCK = 100;

double A[MAX][MAX], B[MAX][MAX], C[MAX][MAX];

void iniciar() {
    srand(42);
    for(int i = 0; i < MAX; i++){
        for(int j = 0; j < MAX; j++){
            A[i][j] = rand() % 100;
            B[i][j] = rand() % 100;
            C[i][j] = 0;
        }
    }
}

void multiplicarMatricesBloques() {
    for(int ii = 0; ii < MAX; ii += BLOCK) {
        for(int jj = 0; jj < MAX; jj += BLOCK) {
            for(int kk = 0; kk < MAX; kk += BLOCK) {
                for(int i = ii; i < min(ii + BLOCK, MAX); i++) {
                    for(int j = jj; j < min(jj + BLOCK, MAX); j++) {
                        for(int k = kk; k < min(kk + BLOCK, MAX); k++) {
                            C[i][j] += A[i][k] * B[k][j];
                        }
                    }
                }
            }
        }
    }
}

int main() {
    iniciar();
    auto start = high_resolution_clock::now();
    multiplicarMatricesBloques();
    auto stop = high_resolution_clock::now();
    auto duration = duration_cast<milliseconds>(stop - start);
    cout << "Tiempo para " << MAX << " dimensiones: "<< duration.count() << "ms" << endl;
}
