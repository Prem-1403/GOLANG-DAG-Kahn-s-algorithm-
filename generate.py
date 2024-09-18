import json
import os
import random
from collections import defaultdict, deque

def generate_vertex_data(vertex_id):
    return {
        "id": vertex_id,
        "data": {
            "name": f"Block {vertex_id}",
            "info": "Generated vertex"
        }
    }

def generate_vertices(n):
    vertices = {}
    for vertex_id in range(n):
        if vertex_id not in vertices:
            vertices[vertex_id] = generate_vertex_data(vertex_id)
    return vertices

def has_cycle_kahn(vertices, edges):
    adj_list = defaultdict(list)
    in_degree = defaultdict(int)
    
    for edge in edges:
        adj_list[edge["src"]].append(edge["dest"])
        in_degree[edge["dest"]] += 1
        if edge["src"] not in in_degree:
            in_degree[edge["src"]] = 0

    zero_in_degree = deque([v for v in vertices if in_degree[v] == 0])
    visited_count = 0

    while zero_in_degree:
        node = zero_in_degree.popleft()
        visited_count += 1
        
        for neighbor in adj_list[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                zero_in_degree.append(neighbor)

    return visited_count != len(vertices)

def generate_edges(vertices, max_edges_per_node):
    edges = []
    added_edges = set()  # To avoid duplicate edges
    vertex_ids = list(vertices.keys())

    # Step 1: Ensure every vertex is connected by creating a backbone chain
    for i in range(len(vertex_ids) - 1):
        src = vertex_ids[i]
        dest = vertex_ids[i + 1]
        edges.append({"src": src, "dest": dest})
        added_edges.add((src, dest))

    # Step 2: Add additional random edges for complexity
    for vertex_id in vertex_ids:
        num_edges = random.randint(1, max_edges_per_node)

        for _ in range(num_edges):
            target_id = random.choice(vertex_ids)
            
            # Ensure we don't create a self-loop or duplicate edges
            if vertex_id != target_id and (vertex_id, target_id) not in added_edges:
                # Ensure we're not creating an edge that introduces a cycle
                if vertex_id < target_id:  # Ensures the graph remains acyclic
                    edges.append({"src": vertex_id, "dest": target_id})
                    added_edges.add((vertex_id, target_id))

    # Check for any cycles, regenerate edges if needed
    if has_cycle_kahn(vertices, edges):
        print("Cycle detected, regenerating edges...")
        return generate_edges(vertices, max_edges_per_node)

    return edges

def save_data_to_file(vertices, edges, filename):
    with open(filename, 'w') as f:
        json.dump({
            "vertices": vertices,
            "edges": edges
        }, f, indent=4)

def main():
    try:
        n = int(input("Enter the number of vertices: "))
        max_edges = int(input("Enter the maximum number of edges per node: "))
        if n < 2 or max_edges < 1:
            print("Number of vertices must be at least 2 and max edges at least 1.")
            return
    except ValueError:
        print("Invalid input. Please enter valid integers.")
        return

    print("Generating vertices...")
    vertices = generate_vertices(n)
    
    print("Generating edges and checking for cycles...")
    edges = generate_edges(vertices, max_edges)
    
    file_path = r'C:\Users\Prem\Desktop\DAG\data\sample_Dag.json'
    print(f"Saving data to file: {file_path}")
    os.makedirs(os.path.dirname(file_path), exist_ok=True)  # Create directory if it does not exist
    save_data_to_file(vertices, edges, file_path)

    print("DAG data saved to 'sample_Dag.json'")

if __name__ == "__main__":
    main()
