#include<bits/stdc++.h>
using namespace std;
class UnionFind {
public:
    vector<int> parent;
    int count; // number of connected components
    
    UnionFind(int n) : count(n), parent(n) {
        for (int i = 0; i < n; ++i)
            parent[i] = i;
    }
    
    int find(int x) {
        if (parent[x] != x)
            parent[x] = find(parent[x]); // path compression
        return parent[x];
    }
    
    void unite(int x, int y) {
        int rootX = find(x);
        int rootY = find(y);
        if (rootX != rootY) {
            parent[rootY] = rootX; // union
            count--;
        }
    }
};

int makeConnected(int n, vector<vector<int>>& connections) {
    if (connections.size() < n - 1) return -1; // Not enough edges to connect all nodes
    
    UnionFind uf(n);
    
    for (const auto& conn : connections) {
        uf.unite(conn[0], conn[1]);
    }
    
    return uf.count - 1; // Need (count-1) additional connections
}
int main(){
 int n = 4;
    vector<vector<int>> connections = {{0, 1}, {0, 2}, {1, 2}};
    
    cout << "Minimum changes required: " << makeConnected(n, connections) << endl;
    
    return 0;
}